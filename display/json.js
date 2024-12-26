const fs = require('fs');
const http = require('http');

// Function to read and parse the file into JSON
function parseFileToJson(filePath) {
    const fileContent = fs.readFileSync(filePath, 'utf-8').split('\n').filter(line => line.trim() !== '');

    const jsonData = [];
    for (let i = 0; i < fileContent.length; i++) {
        const [id, username, bookName, author, imageUrl] = fileContent[i].split(' ');

        jsonData.push({
            username: username.trim(),
            id: id.trim(),
            bookName: bookName.trim(),
            author: author.trim(),
            imageUrl: imageUrl.trim()
        });
    }

    return jsonData;
}

// HTTP server to serve JSON data
const server = http.createServer((req, res) => {
    if (req.url === '/anxielarchives/community' && req.method === 'GET') {
        try {
            const inputFilePath = '../data/books.txt';
            const jsonData = parseFileToJson(inputFilePath);

            res.writeHead(200, { 'Content-Type': 'application/json' });
            res.end(JSON.stringify(jsonData, null, 2));
        } catch (error) {
            res.writeHead(500, { 'Content-Type': 'text/plain' });
            res.end('Error reading or processing the file: ' + error.message);
        }
    } else {
        res.writeHead(404, { 'Content-Type': 'text/plain' });
        res.end('Not Found');
    }
});

// Start the server
const PORT = 10000;
server.listen(PORT, () => {
    console.log(`Server is running at http://localhost:${PORT}/anxielarchives/community`);
});
