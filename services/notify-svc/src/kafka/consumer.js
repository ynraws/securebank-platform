const kafka = require('kafka-node');

const KAFKA_HOST = process.env.KAFKA_BOOTSTRAP_SERVERS || 'localhost:9092';

function startConsumer() {
  try {
    const client = new kafka.KafkaClient({ kafkaHost: KAFKA_HOST });
    const consumer = new kafka.Consumer(
      client,
      [{ topic: 'payment-events', partition: 0 }],
      { autoCommit: true, groupId: 'notify-group' }
    );

    consumer.on('message', (message) => {
      try {
        const event = JSON.parse(message.value);
        const { eventType, paymentId, fromAccountId, amount, currency } = event;

        console.log(`\n📨 Event received: ${eventType}`);

        if (eventType === 'PAYMENT_INITIATED') {
          console.log(`📧 SMS: Dear customer, payment of ${currency} ${amount} initiated. Ref: ${paymentId}`);
          console.log(`📧 EMAIL: Payment initiated for account ${fromAccountId}`);
        } else if (eventType === 'PAYMENT_COMPLETED') {
          console.log(`📧 SMS: Payment of ${currency} ${amount} completed successfully!`);
        } else if (eventType === 'PAYMENT_FAILED') {
          console.log(`🚨 ALERT: Payment ${paymentId} FAILED for account ${fromAccountId}`);
        }
      } catch (err) {
        console.error('Error processing message:', err.message);
      }
    });

    consumer.on('error', (err) => {
      console.error('Kafka consumer error:', err.message);
      setTimeout(startConsumer, 5000);
    });

    console.log(`Kafka consumer connected to ${KAFKA_HOST}`);
  } catch (err) {
    console.error('Failed to connect to Kafka:', err.message);
    setTimeout(startConsumer, 5000);
  }
}

startConsumer();
