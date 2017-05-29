////////////////
//  thermoIP  //
////////////////
#include <SPI.h>
#include <Ethernet2.h>

const int tmp36pin = 0;
byte mac[] = {
  0xDE, 0xAD, 0xBE, 0xEF, 0xFE, 0xED
};
IPAddress ip(192, 168, 1, 177);

EthernetServer server(80);

void setup() {
  Serial.begin(9600);

  Ethernet.begin(mac, ip);
  Serial.println("init");

  delay(1000);
  Serial.println("Ready");
  server.begin();
  Serial.print("Server is at ");
  Serial.println(Ethernet.localIP());
}

void loop() {
  // listen for incoming clients
  EthernetClient client = server.available();
  if (client) {
    Serial.println("new client");

    boolean currentLineIsBlank = true;
    while (client.connected()) {
      if (client.available()) {
        char c = client.read();
        Serial.write(c);
	// if EOL then send a reply
        if (c == '\n' && currentLineIsBlank) {
          sendTemp(client);          
          break;
        }
        if (c == '\n') {
          // you're starting a new line
          currentLineIsBlank = true;
        } else if (c != '\r') {
          // you've gotten a character on the current line
          currentLineIsBlank = false;
        }
      }
    }
    delay(1);
    client.stop();
  }
}

void sendTemp(EthernetClient client) {
  int reading = analogRead(tmp36pin);

  float voltage = (reading * 5.0) / 1024;
  float tempC = (voltage - 0.5) * 100;
  String temp = String(tempC);

  Serial.println(temp);

  // send a standard http response header
  client.println("HTTP/1.1 200 OK");
  client.println("Content-Type: application/json");
  client.println("Connection: close");  // the connection will be closed after completion of the response
  //client.println("Refresh: 5");  // refresh the page automatically every 5 sec
  client.println();

  // send JSON format response
  client.print("{ \"value\": ");
  client.print(temp);
  client.print(", \"unit\": \"celsius\"}");
}
