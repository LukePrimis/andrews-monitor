# Andrews Dining Hall Monitor

A Go program to monitor when times were available to order for Brown's Andrews dining hall. Used during the portion of the pandemic when the dining hall was only available for online order.

## A Few Notes

Just a couple of things to note!

### How to use it

- As will be mentioned before, this bot is obsolete and just for educational purposes. But, all it would require to run is the user to input their info into `config/config.go`
- It takes Twilio API info for the SMS messaging as well as a Discord Webhook URL (in addition to the username/password to log in)

### What could be done better

- For the parsing, I just wrote a really janky parsing helper function because I did this all in one afternoon. A more elegant solution might be using the GoQuery package for JQuery parsing, or possibly RegEx
- I didn't feel like messing around with GoQuery as it would add more parsing time compared to the janky string parsing I did, and I didn't want to deal with RegEx for such a silly and quick project
- Honestly, you might not need all the headers for DUO login but in the bot business, I always just copy the same exact headers that a manual user would have when using Chrome. I always find it better to blend in
- On that same note, when I compile on my local computer this bot also used my modified Go HTTP client that had spoofed TLS ClientHello packets as well as modified HTTP2 code to appear _very_ similar to Google Chrome's request. Again, this was far overkill and if you don't know what that means then don't worry because it's entirely irrelevant to the code of this bot/monitor
- There's also just some generally sloppy code. For example, I just assume there will be a `Location` header on a status 302 redirect when in reality there is no guarantee of one!

### Disclaimers

- This monitor is no longer useful as the cafeteria has returned to normal use!
- I have not maintained the DUO login funcionality, so if you are looking for a bot that can log into DUO this repo's code may not work anymore as they frequently change stuff
- The reverse engineering of the DUO login flow is **purely** for webscraping and educational purposes. Do **NOT** use this login code for any malicious purposes. It would be pointless to even try and do so considering it still requires the manual app 2FA of whichever user has their info input into the bot!
- There is no private information in here at all. All URLs for the API endpoints are publicly available in the Javascript code of DUO and on the CBORD login page (I just grabbed them from Chrome DevTools network logging), so there's no funny business here! All this bot did was login to CBORD, I authorized the DUO login from my phone, and it sent a Discord/SMS message to me when I could order food!

### Why I did this

- I could've just set reminders to tell me to order food because the time slots went live almost always at the same time, but I wanted to make this bot to post publicly eventually
- All of the work I do for bots like Akari is always paid and the code is not available to be published publicly, but I wanted to really quickly throw together an example bot as a lot of the methods for putting together request payloads, parsing from the page, logging in, etc. is _very_ similar to the exact stuff in the bots I work on
- In the same vein as the last bulletpoint, the only difference between this bot and the private bot I work for is what websites it's designed to work on, the real bot I work on has a frontend to control it over GRPC, and this bot also runs single-threaded but if you wanted to run a bunch of "tasks" you could just use Goroutines and make it multithreaded!
- Another difference is the info required for each task. For this bot, the relevant data is the login info, Discord Webhook URL, etc. But, a typical shoe-bot or retail-bot task would need stuff like payment info, shipping address, email, etc.
