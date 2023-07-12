# Telegram Quiz Bot

---------------

## Introduction

This projects implements Telegram bot which have its own dataset
of questions on different themes and provides users an interface
to set up configuration for quiz (Sphere, Section, Difficulty, Amount of questions)
and then by turns sends polls with questions and options to a user.

Dataset of questions is generated with ChatGPT and serves as an example
to show Bot logic.

Try Bot: https://t.me/universal_quiz_bot

------------------
## Technologies
- #### Go 1.18
- #### Docker / Docker Compose
- #### PostgreSQL
- #### MongoDB 
- #### CI/CD (GitHub Actions)

-----------------

## Launch 

To launch this bot locally inside of docker containers you should:

- Set up MongoDB on your computer or Mongo Atlas, also notice
that names of the collections and database in config and your DB 
should match, don't forget to change connection string 
if you use Mongo Atlas.
- Make your own Telegram Bot with BotFather and get its token.
- Enter your values to environment variables in docker-compose file.
- Check if Docker and Docker Compose are installed on your device.
- Run command ```make run```

----------------

