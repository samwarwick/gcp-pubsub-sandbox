const {PubSub} = require('@google-cloud/pubsub');

async function publish(projectId, topicName, data) {
    const pubsub = new PubSub({projectId});
    const topic = pubsub.topic(topicName);

    // https://stackoverflow.com/questions/59758539/what-is-the-correct-way-to-publish-to-gcp-pubsub-from-a-cloud-function
    topic.publish(Buffer.from(data));
}

var subject = 'Hello world!'
if (process.argv.length > 2 && process.argv[2] !== "") {
    subject = process.argv[2];
}

const data = JSON.stringify({
    'origin': 'node',
     subject
});

publish('gps-demo', 'demo-topic', data);
