using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading;
using System.Threading.Tasks;

namespace ReadGmail
{
    internal class Program
    {
       
        static void Main(string[] args)
        {
            string username = "example@gmail.com";
            string password = "passsword";
            MailRepository mailBox;
            Console.WriteLine("Connecting MailBox....");
            //connect Gmail
            mailBox = new MailRepository(
                       "imap.gmail.com",
                       993,
                       true,
                       username,
                       password
                   );
            Console.WriteLine("Connected MailBox");
            var emailList = mailBox.GetAllMails("inbox");            
            foreach(var email in emailList)
            {
                Console.WriteLine(email.From.Email,email.Subject);
            }
        }
    }
}
