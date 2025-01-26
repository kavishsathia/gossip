<div align="center">
  <img width="550" alt="Screenshot 2025-01-07 at 20 43 32" src="https://github.com/user-attachments/assets/292a93b0-6ea9-4680-9f43-0d77b4cba63a" />
  <br><br>
  
  [![Netlify Status](https://api.netlify.com/api/v1/badges/f8790419-0f6c-4023-baf0-ac9bb15e8179/deploy-status)](https://app.netlify.com/sites/uniconnweb/deploys)
</div>

<div align="center">
<h1>Uniconn: My Implementation of a Simple Web Forum Application</h1>

**Author:** Kavish Sathia  
**Date:** 3rd January 2025

</div>

## For Quick Access:
- <a href="https://uniconnweb.netlify.app/">The Website</a>
- <a href="https://gossip.s6wyfaw6z9q0r.ap-southeast-1.cs.amazonlightsail.com/swagger/index.html">Swagger API Documentation</a>
- [Developer Manual](#developer-manual)
- [User Manual](#user-manual)


<br>

## Table of Contents

- [Features](#features)
  - [Planned or In-Progress Features](#planned-or-in-progress-features)
- [Tech Stack](#tech-stack)
  - [Databases](#databases)
- [User Interface](#user-interface)
- [Data Model](#data-model)
- [Architecture](#architecture)
- [API Documentation](#api-documentation)
- [Developer Manual](#developer-manual)
- [User Manual](#user-manual)
- [Accomplishments](#accomplishments)

<br>

## Features

These are the features that I have included in Uniconn. To try out these features, go to: [Uniconn Web](https://uniconnweb.netlify.app) (You can click on "Lurk" to skip the authentication process):

- Account-based authentication with JWT stored in an HTTP Cookie.
- Write threads using markdown and include a thumbnail picture for more engagement. Threads are editable and deletable.
- Real-time notifications when the user's thread or comment is liked or commented on.
- Nested comments.
- Mobile responsiveness.
- Progressive web app (allowing users to install the web app onto their devices).
- Searching for forums based on title, description, tags, and author username. The response is paginated (click on "Load more" at the bottom of the sidebar to load more threads).
- A point system (Aura) that will increase when a user's thread or comment gets liked, incentivizing participation.
- Content moderation using AI and community feedback.
- AI-generated descriptions for threads.
- Fact checking for threads

### Planned or In-Progress Features

- Communities feature with each community having its own theme and color scheme.
- Include tests in future versions

<br>

## Tech Stack

The frameworks and languages used follow the assignment guidelines, using React and TypeScript for the frontend and Golang with the Gin web framework for the backend. Additional libraries and their purposes are listed below:

- **React Router**: For frontend navigation and handling different routes.
- **TailwindCSS**: For inline styling.
- **MUI**: Used by CVWO, so I tried it out.
- **MDXEditor**: For the markdown editor.
- **Lucide-React**: For iconography.
- **Gin**: For the web framework.
- **GORM**: Used as the ORM to query the database.
- **Redis-Go**: To maintain PubSub connections to the Redis server.

### Databases

I used two databases for this assignment:

1. **PostgreSQL**: To store application data. Threads, comments, and user data are stored here.
2. **Redis**: As the PubSub server to publish and listen to real-time notifications from user activity.

<br>

## User Interface

There are 4 interfaces for larger screens and 5 interfaces for smaller screens. These consist of:

- Login interface.
- Sign up interface.
- Thread interface (this includes the comments).
- Thread editor interface.
- Thread list interface (exists as a sidebar on larger screens but as its own interface on mobile screens).

<br>

## Data Model

<div align="center">
<img width="1028" alt="data_model(1)" src="https://github.com/user-attachments/assets/b1f907d0-40da-4605-b84c-629b7c01fc2b" />

_Figure: The ERD diagram for Uniconn's relational database._

</div>

<br>

## Architecture

<div align="center">
<img width="920" alt="architecture(1)" src="https://github.com/user-attachments/assets/8ca0f41f-2284-4327-ac71-c124f5c9ad5a" />

_Figure: The tech architecture for Uniconn._

</div>

The architecture consists of 3 segments: the frontend which is hosted on Netlify,
the backend which is hosted on AWS Lightsail as a Docker container and the data
layer which has two databases, PostgreSQL and Redis. The Redis database is not used
for data storage but for its PubSub capabilities. I will now outline how the real time
notification system works: For each user, there exists a channel on the Redis server whose
name is the user ID of the user to which it corresponds. Every time an action has been
taken on a user's post or comment, a message will be published to the channel that
corresponds to the user. If the user is connected to the WebSocket endpoint on the Gin
backend, the WebSocket endpoint will read the published message and send it to the user
through the Websocket connection. When the frontend receives the message, it will display a notification.

As for the moderation feature, I used the free service from OpenAI
to moderate posts based on a few categories. The description and the fact checking is done using the
GPT-4o-mini model from OpenAI.

<br>

## API documentation

<a href="https://gossip.s6wyfaw6z9q0r.ap-southeast-1.cs.amazonlightsail.com/swagger/index.html">Go to the Swagger documentation generated with OpenAPI specs.</a>

<br>

## Developer Manual

1. Clone this repository. Run:

```console
git clone git@github.com:kavishsathia/gossip.git
```

2. Change your directory to the backend. Run:

```console
cd gossip/backend
```

3. Download the required modules. Run:

```console
go mod download
```

4. Copy the text in the .env.example file into your .env file for both your frontend and backend.

5. Start a Postgres server locally or on NeonDB (neon.tech). Start a redis-server locally by running the redis-server command. If not found, download the CLI.

6. Fill up the .env file with the relevant environment variables. You can get your OpenAI API key from the Open AI console.

7. Start the server. Run:

```console
go run .
```

8. Explore the <a href="https://gossip.s6wyfaw6z9q0r.ap-southeast-1.cs.amazonlightsail.com/swagger/index.html">API documentation</a>. to learn more.
  
9. Start a new terminal session.

10. Change your directory to the frontend. Run:

```console
cd gossip/frontend
```

11. Download the required modules. Run:

```console
npm install
```

12. Start the server. Run:

```console
npm run dev
```

13. The server will start at localhost:5173.

<br>

## User Manual

1. How to create a new account? Click on sign up on the login page.
2. Can I view posts without signing up? Yes, you can! Just click on ”Lurk” on the login page, however
   if you try to like or comment, you will be redirected to the login page.
3. How to search for posts? Use the search bar on the left sidebar. The response is paginated, so click on "Load more" at the bottom of the sidebar to load more threads.
4. How to search by tags and author username? Add ”#” in front of tags and ”@” in front of
   usernames when searching in the search bar. For example, to find all posts written by me with the
   ”nus” tag and contains the word ”CS”, you can search ”@kavish #nus cs”.
5. Is there a quicker way to search by tags and username? Yes, you can click on the tags in the posts
   and click on the author username in the posts (then you will be shown a modal in which ”View
   their Threads” is an option”)
6. How to view a thread? Click on any of the threads in the left sidebar.
7. How to view comments? Scroll down and you will see comments if there are any. You can view
   nested comments by clicking on ”Show replies”, and you will also be presented with text input to
   reply to that comment.
8. How to write comments? Simply fill in the text input at the bottom of the thread and Enter to
   post the comment.
9. How to like threads and comments? Click on the heart button, if you are being redirected to the
   login page, it means you are lurking.
10. How to create my own thread? Click on the ”Plus” button at the bottom right of the page. An
    editor page will open up. You can submit a URL for the thumbnail picture. You can change the
    title by editing the ”Untitled Thread” text. You can also add more tags by clicking on ”Add tag”
    or the ”X” button next to the tags. The body of the post can be edited within the markdown
    editor. After you’re done, click on the tick button on the bottom right corner.
11. How do I edit my thread or comment? For your own threads and comments, you will see a pencil
    icon on the top right corner. Click on that to edit. For threads, the familiar markdown editor
    will open up, and after you’re done, click on the tick button. For comments, an ”Edit” button
    will replace the pencil icon on the top right corner, and the comment text will be replaced with a
    textarea. Simply edit the text in the textarea and click on the ”Edit” button.
12. How to delete comments and threads? Click on the trash icon next to the pencil icon.
13. How to report threads? Click on the ”Report” button if it appears on the top of the thread. It
    will only appear if the thread has been flagged by the moderation service.

<br />

## Accomplishments

This is my first (almost) complete open-source project on GitHub, and it has been a game-changer for
me as a programmer. Through this project, I learned so much that I genuinely feel like I’ve leveled up.
One of the biggest challenges was setting up the markdown editor and deploying the application as a
container on AWS Lightsail (tasks that initially took me around 5 hours and a full day respectively).
More recently though, when I had to do the same for Hack&Roll 2025 (a hackathon), I broke my personal
best, configuring the Markdown editor in just 10 minutes and deploying to AWS Lightsail in 15 minutes.
These improvements contributed vastly to my team winning at the hackathon, and I owe a huge thanks
to the CVWO team for helping me hone my skills along the way.
