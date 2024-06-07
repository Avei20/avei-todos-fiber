## Avei's Todo's Backend Clean Architecture

Hi, this is my side project for handling my personal task. I create this to maximize my tasking on different project to keep all of my project stay in the line.

The idea is simple, create a todo but every todos can have child and parent. Then they have a separated project. If i am not lazy, i will integrate this with google calendar api. So, the task will updated to google calendar.

But, the main reason i built this is to managing my work and personal task. Espocially when i need to make a monthly report for my task. (I know i can use Jira, but just an over engineer xD).

## Tech Stack
Golang, Clean Architecure, Fiber, NO ORM BULLSHIT RAW QUERY IS THE WAY using pgx, Docker, Postgresql on Neon.

## How to run
```
go run main.go
```

or build the docker images.

### Deployment
This application is deployed on Google Cloud Run. (All hail Cloud Run free tier. I Love U Google)
