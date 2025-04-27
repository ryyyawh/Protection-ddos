DDoS Protection and Attack Scripts

This repository contains two components:

1. Golang DDoS Protection Server: A server with multiple layers of protection against DDoS attacks.


2. Python DDoS Attack Script: A simple script to simulate DDoS attacks.

Table of Contents

Golang DDoS Protection Server

Installation

Usage

How It Works


Golang DDoS Protection Server

A DDoS protection server written in Golang that uses multiple techniques to prevent attacks such as rate limiting, IP blocking, CAPTCHA challenges, and geofencing.

Installation

1. Install Go:
If you haven't installed Go yet, you can download it from the official Go website: https://golang.org/dl/


2. Clone the repository:

git clone https://github.com/ryyyawh/Protection-ddos.git
cd Protection-ddos


3. Run the Go server:

go run Protection_ddos.go



Usage

The server will start on port 8080 by default. You can access it in your browser or through curl:

curl http://localhost:8080

Rate Limiting: The server allows only 2 requests per 5 seconds from the same IP address. Exceeding this limit will result in a 429 Too Many Requests response.

IP Blocking: Blacklisted IPs (e.g., 192.168.1.10) will receive a 403 Forbidden response.

Whitelisting: Only IPs in the whitelist (e.g., 127.0.0.1) can access the server.

Geofencing: Only IPs from the US region are allowed to access the server. Requests from other regions will get a 403 Forbidden response.

Captcha: For suspicious traffic (e.g., POST requests), users will be required to complete a CAPTCHA.


How It Works

Rate Limiting: The server limits the number of requests from the same IP address to 2 requests every 5 seconds using the golang.org/x/time/rate package.

IP Blacklisting: Blacklisted IPs are blocked from accessing the server.

IP Whitelisting: Whitelisted IPs can access the server without any restriction.

Geofencing: The server blocks access from IPs outside the US.

Captcha: For POST requests (commonly used in HTTP floods), the server will challenge users to complete a CAPTCHA.

©® Xylays Developer
Contact me : t.me/conquerryy
