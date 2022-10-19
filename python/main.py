import imaplib
import email

server = "imap.gmail.com"
imap = imaplib.IMAP4_SSL(server)

username = "example@gmail.com"
password = "password"

imap.login(username, password)
res, messages = imap.select('Inbox')

# get whole emails count
messages = int(messages[0])

# set emails count to read
n = 3

for i in range(messages, messages - n, -1):
    res, msg = imap.fetch(str(i), "(RFC822)")
    for response in msg:
        if isinstance(response, tuple):
            msg = email.message_from_bytes(response[1])
            From = msg["From"]
            subject = msg["Subject"]
            print("From : ", From)
            print("subject : ", subject)
