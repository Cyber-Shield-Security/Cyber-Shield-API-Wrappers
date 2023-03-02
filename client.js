const net = require('net');
const apiKey = '4IqF7qAxP9JObVr5R16XOxpKePKPMAFpaZ1gft7Ez75IVp1H';
const appVersion = 'APPLICATION_NAME_AND_VERSION';
const csBackend = '15.204.58.133';

let licenseID = {
    auth_status: false,
    id: apiKey,
    username: '',
    discord_id: '',
    rank: ''
};

const client = net.createConnection({
    host: csBackend,
    port: 17234
}, () => {
    // Connection successful, send authentication data
    client.write(`${appVersion}\n${JSON.stringify(licenseID)}\n`);
});

client.on('data', (data) => {
    const message = data.toString().trim();
    if (message === 'close') {
        console.log('[!] Error, The owner has disconnected you from using Cyber Shield...');
        client.end();
    } else if (message === 'banned') {
        console.log('[!] Error, You have been banned from using Cyber Shield...');
        client.end();
    } else if (message.startsWith('{') && message.endsWith('}')) {
        const jsonData = JSON.parse(message);
        licenseID.auth_status = jsonData.status;
        licenseID.username = jsonData.discord_tag;
        licenseID.discord_id = jsonData.discord_id;
        licenseID.rank = jsonData.rank;
        console.log(`Successfully authed! ${licenseID.rank}`);
        // start sending or receiving data with server here
    }
});

client.on('end', () => {
    // Connection ended
    console.log('Disconnected from server');
});

client.on('error', (err) => {
    console.log(`[!] Error: ${err.message}`);
});

// github.com/9xN
