import { Construct } from 'constructs';
import { App, TerraformStack } from 'cdktf';
import { GoogleProvider, Project, CloudSchedulerJob } from './.gen/providers/google';

class MyCloudSchedulerStack extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);

    new GoogleProvider(this, 'google', {
      project: 'YOUR_GCP_PROJECT_ID',
      region: 'us-central1', // Change this to your desired region
    });

    const project = new Project(this, 'project', {
      projectId: 'YOUR_GCP_PROJECT_ID',
    });

    new CloudSchedulerJob(this, 'my-scheduled-job', {
      name: 'my-scheduled-job',
      description: 'My Cloud Scheduler Job',
      schedule: '*/5 * * * *', // Change this to your desired schedule (cron expression)
      timeZone: 'UTC', // Change this to your desired time zone
      pubsubTarget: [
        {
          topicName: 'YOUR_PUBSUB_TOPIC', // Replace this with your existing Pub/Sub topic
        },
      ],
    });
  }
}

const app = new App();
new MyCloudSchedulerStack(app, 'MyCloudSchedulerStack');
app.synth();

