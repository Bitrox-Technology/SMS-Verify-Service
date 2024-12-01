# SMS Verification Service

A Go-based SMS verification service leveraging the [Twilio API](https://www.twilio.com/) and the Gin web framework. This service sends verification codes via SMS and validates user input against the code sent.

---

## 🚀 Features

- Send SMS verification codes to user phone numbers.
- Validate user-provided verification codes.
- Integration with Twilio's Verify API.

---

## 📋 Prerequisites

1. **Twilio Account**: You need a Twilio account to obtain the required credentials:
   - `TWILIO_ACCOUNT_SID`
   - `TWILIO_AUTH_TOKEN`
   - `TWILIO_PHONE_NUMBER` (Service SID for the Verify API)
2. **Go**: Version 1.19 or later.

---

## 🛠 Setup Instructions

### 1. Clone the Repository

```bash
git clone https://github.com/Bitrox-Technology/SMS-Verify-Service.git
cd sms-verify-service

.env file

TWILIO_ACCOUNT_SID=<your-account-sid>
TWILIO_AUTH_TOKEN=<your-auth-token>
TWILIO_PHONE_NUMBER=<your-service-id>

run server
go run cmd/main.go