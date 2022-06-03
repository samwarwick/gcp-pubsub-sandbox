const {PubSub} = require('@google-cloud/pubsub');

async function subscribe(projectId, topicName, subscriptionName) {
    const pubsub = new PubSub({projectId});
    const topic = pubsub.topic(topicName);
    const subscription = topic.subscription(subscriptionName);

    subscription.on('message', async (message) => {
        const json = JSON.parse(message.data.toString());

        console.log(`Received message (id=${message.id}):`, json);
        console.log('Subject:', json.subject);

        message.ack(); // void. clears message from queue

      // https://stackoverflow.com/questions/65702233/delayed-acknowledging-gcloud-pub-sub-message
      // Do not quit app immediately after ack otherwise it might not complete!
        if (json.subject === 'quit') {
            console.log("Quitting...")
            await new Promise(resolve => setTimeout(resolve, 5000))
            process.exit(0);
        }    
    });

    subscription.on('error', error => {
        console.error('ERROR:', error);
        process.exit(1);
    });
}

console.log("Node subscriber")
subscribe('gps-demo', 'demo-topic', 'demo-sub'); 
