const express = require('express');
const app = express();
app.use(express.json());

const PORT = process.env.PORT || 8084;

app.get('/health', (req, res) => {
  res.json({ status: 'healthy', service: 'notify-svc' });
});

app.post('/api/v1/notify', (req, res) => {
  const { type, accountId, message, amount } = req.body;
  const notification = {
    id: Date.now().toString(),
    type,
    accountId,
    message,
    amount,
    timestamp: new Date().toISOString(),
    delivered: true
  };
  console.log(`📧 NOTIFICATION [${type}]: Account ${accountId} - ${message}`);
  res.status(201).json(notification);
});

app.listen(PORT, () => {
  console.log(`notify-svc running on port ${PORT}`);
  require('./kafka/consumer');
});
