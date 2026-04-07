const express = require('express');
const app = express();
app.use(express.json());

// POST API
app.post('/predict', (req, res) => {
    const { x } = req.body;
    const y = 2 * x + 3;
    res.json({ prediction: y });
});

// GET for browser
app.get('/', (req, res) => {
    res.send("Service is running");
});

app.get('/predict', (req, res) => {
    res.send("Use POST request with JSON {x:value}");
});

// START SERVER (ALWAYS LAST)
app.listen(3000, () => {
    console.log("Server running on port 3000");
});
