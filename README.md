# Ghost Without Guilt 👻💬

### Ever felt guilty for not replying… but just didn’t feel like it?  
Or want to ghost someone without them ever noticing?

Meet **Ghost Without Guilt** — your new chat sidekick!

This project reads your WhatsApp messages, sends them to ChatGPT, and magically crafts replies that sound like *you* — only cooler, wittier, and way less stressed.

---

### Why suffer from awkward silences when you can:  
- Skip the guilt of leaving someone on read  
- Ghost like a pro without awkward “sorry, was busy” excuses  
- Send automatic replies that actually make sense  

Say goodbye to texting anxiety and hello to effortless ghosting with style!


# Upcoming v1.0.0-beta Feature List: WhatsApp + ChatGPT Integration

## 1. Receive Messages from WhatsApp
- Connect to WhatsApp via API WhatsApp Business API
- Capture incoming user messages in real time.
- Extract message content and metadata:
  - Sender ID
  - Timestamp
  - Message type (text, media, etc.)

## 2. Send Message to ChatGPT
- Forward the received message to the ChatGPT API.
- Format the message to maintain context and clarity.
- Optionally manage user sessions for conversational continuity.

## 3. Get Response from ChatGPT
- Receive response from ChatGPT.
- Parse and validate the content:
  - Sanitize special characters
  - Handle message length limits
  - Ensure proper formatting for WhatsApp

## 4. Send Response Back to WhatsApp
- Format ChatGPT’s reply for WhatsApp delivery.
- Send the response to the original

