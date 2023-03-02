<?php
$host = '15.204.58.133';
$port = 17234;
$apiKey = '4IqF7qAxP9JObVr5R16XOxpKePKPMAFpaZ1gft7Ez75IVp1H';
$appVersion = 'APPLICATION_NAME_AND_VERSION';

$licenseID = array(
    'auth_status' => false,
    'id' => $apiKey,
    'username' => '',
    'discord_id' => '',
    'rank' => ''
);

$client = stream_socket_client("tcp://$host:$port", $errno, $errstr);

if (!$client) {
    echo "[!] Error: $errstr ($errno)\n";
} else {
    // Connection successful, send authentication data
    fwrite($client, "$appVersion\n" . json_encode($licenseID) . "\n");
}

while ($data = fgets($client)) {
    $message = trim($data);
    if ($message === 'close') {
        echo '[!] Error, The owner has disconnected you from using Cyber Shield...';
        fclose($client);
        break;
    } else if ($message === 'banned') {
        echo '[!] Error, You have been banned from using Cyber Shield...';
        fclose($client);
        break;
    } else if (strpos($message, '{') === 0 && strrpos($message, '}') === strlen($message) - 1) {
        $jsonData = json_decode($message, true);
        $licenseID['auth_status'] = $jsonData['status'];
        $licenseID['username'] = $jsonData['discord_tag'];
        $licenseID['discord_id'] = $jsonData['discord_id'];
        $licenseID['rank'] = $jsonData['rank'];
        echo "Successfully authed! {$licenseID['rank']}";
        // start sending or receiving data with server here
    }
}

fclose($client);
echo "Disconnected from server";



// github.com/9xN

?>
