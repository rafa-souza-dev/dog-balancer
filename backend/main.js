const http = require('http');

const PORT = process.env.PORT || 8001;

const server = http.createServer((req, res) => {
  if (req.method === 'GET' && req.url === '/helthcheck') {
    const response = JSON.stringify({ message: 'OK' });

    res.writeHead(200, {
      'Content-Type': 'application/json',
      'Content-Length': Buffer.byteLength(response),
    });

    res.end(response);
  } else {
    res.writeHead(404, { 'Content-Type': 'application/json' });
    res.end(JSON.stringify({ message: 'Not Found' }));
  }
});

server.listen(PORT, () => {
  console.log(`HTTP Server running at http://localhost:${PORT}`);
});
