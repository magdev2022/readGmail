const Imap = require('imap');
const { simpleParser } = require('mailparser');
const username = "jin1991919@gmail.com"
const password = "mqspwumvjpybzehy"
var imapConfig = {
    user: username,
    password: password,
    host: 'imap.gmail.com',
    port: 993,
    authTimeout: 15000,
    tls: true,
    tlsOptions: { rejectUnauthorized: false }
  };
  const imap = new Imap(imapConfig);
  imap.once('ready', () => {
    imap.openBox('INBOX', false, () => {
      imap.search(['UNSEEN'], (err, results) => {
        const f = imap.fetch(results, { bodies: '' });
        f.on('message', msg => {
          msg.on('body', stream => {
            simpleParser(stream, async (err, parsed) => {
              const { from, subject, textAsHtml, text } = parsed;              
              console.log(subject);
            });
          });
        });
        f.once('error', ex => {
          return Promise.reject(ex);
        });
        f.once('end', () => {
          console.log('Done fetching all messages!');
          imap.end();
        });
      });
    });
  });

  imap.once('error', err => {
    console.log(err);
  });

  imap.once('end', () => {
    console.log('Connection ended');
  });
  imap.connect();