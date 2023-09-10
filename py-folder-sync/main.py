import argparse
import time
import os
import hashlib
import shutil
import logging

# Service that has following goals
# Sync between two directories uni-directionally
# New files and folders in source are replicated
# Files modified are replicated
# Files and folders that are deleted are also deleted from replica
# Handle arbitrary levels of nesting in folders

def setup_logger(log_file_path: str):
    global logger
    logger = logging.getLogger('sync-service')
    logger.setLevel(logging.DEBUG)
    file_handler = logging.FileHandler(log_file_path)
    file_handler.setLevel(logging.DEBUG)
    console_handler = logging.StreamHandler()
    console_handler.setLevel(logging.DEBUG)

    formatter = logging.Formatter('%(asctime)s - %(name)s - %(levelname)s - %(message)s')
    file_handler.setFormatter(formatter)
    console_handler.setFormatter(formatter)

    logger.addHandler(file_handler)
    logger.addHandler(console_handler)

def load_folder_state(folder_path: str):
    if not os.path.exists(folder_path):
        logger.error(f'Folder Path: {folder_path} does not exist')
        exit(1)

    files = os.listdir(folder_path)
    files = [file for file in files if os.path.isfile(os.path.join(folder_path, file))]
    
    state = { 'files': {}, 'subdir': {}}
    for root, subdir_list, files in os.walk(folder_path):
        # Setup sub-directory in state, can be opti
        for subdir in subdir_list:
            subdir = os.path.relpath(os.path.join(root, subdir), folder_path)
            state['subdir'][subdir] = True
        for file in files:
            file_path = os.path.join(root, file)
            with open(file_path, 'rb') as f:
                md5 = hashlib.md5(f.read()).hexdigest()
                relative_path = os.path.relpath(file_path, folder_path)
                state['files'][relative_path] = md5

    return state

def sync_folder(src_folder, replica_folder):
    src = load_folder_state(src_folder)
    dest = load_folder_state(replica_folder)

    for src_file, src_file_hash in src['files'].items():
        file_dir = os.path.dirname(src_file)
        # file is not in replica
        if not src_file in dest['files']:
            logger.info(f'Copying New File {src_file} #{src_file_hash}')
            shutil.copytree(os.path.join(src_folder, file_dir), os.path.join(replica_folder, file_dir), dirs_exist_ok=True)
        # file modified
        elif src_file_hash != dest['files'][src_file]:
            logger.info(f'Overwriting Modified File {src_file} #{src_file_hash}')
            shutil.copytree(os.path.join(src_folder, file_dir), os.path.join(replica_folder, file_dir), dirs_exist_ok=True)
    for replica_file, replica_file_hash in dest['files'].items():
        # file removed, since replica is described here to be exact, this will 
        # also handle case when file was not replicated by this service, 
        # but is preset in the replica, they will be deleted
        if not replica_file in src['files']:
            logger.info(f'Deleting File {replica_file} #{replica_file_hash}')
            os.remove(os.path.join(replica_folder, replica_file))

    # Handle the case when a new directory is created without any files
    for sub_dir, _ in src['subdir'].items():
        if not sub_dir in dest['subdir']:
            logger.info(f'Creating Empty Directory {sub_dir}')
            shutil.copytree(os.path.join(src_folder, sub_dir), os.path.join(replica_folder, sub_dir), dirs_exist_ok=True)
    # Handle the case when a directory is deleted
    for sub_dir, _ in dest['subdir'].items():
        # file removed, since replica is described here to be exact, this will 
        # also handle case when file was not replicated by this service, 
        # but is preset in the replica, they will be deleted
        if not sub_dir in src['subdir']:
            logger.info(f'Deleting Directory {sub_dir}')
            shutil.rmtree(os.path.join(replica_folder, sub_dir), ignore_errors=True) # ignore errors when parent dir is already deleted, so this doesn't exist

if __name__ == '__main__':
    parser = argparse.ArgumentParser(description="Sync folders daemon")
    parser.add_argument('--src', help="Src folder path", required=True)
    parser.add_argument('--dest', help="Destination folder path", required=True)
    parser.add_argument('--log-file', required=False, default='.sync-log')
    parser.add_argument('--interval', help="Sync interval. Default 5 seconds", required=False, default=5)

    args = parser.parse_args()

    setup_logger(log_file_path=args.log_file)
    while True:
        sync_folder(src_folder=args.src, replica_folder=args.dest)
        time.sleep(args.interval)