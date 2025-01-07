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

<br>

## Table of Contents

- [Features](#features)
  - [Planned or In-Progress Features](#planned-or-in-progress-features)
- [Tech Stack](#tech-stack)
  - [Databases](#databases)
- [User Interface](#user-interface)
- [Data Model](#data-model)
- [Architecture](#architecture)
  - [Real-Time Notification System](#real-time-notification-system)
- [Challenges](#challenges)
- [Developer Manual](#developer-manual)
- [Appreciation](#appreciation)

<br>

## Features

These are the features that I have included in Uniconn. To try out these features, go to: [Uniconn Web](https://uniconnweb.netlify.app) (You can click on "Lurk" to skip the authentication process):

- Account-based authentication with JWT stored in an HTTP Cookie.
- Write threads using markdown and include a thumbnail picture for more engagement.
- Real-time notifications when the user's thread or comment is liked or commented on.
- Nested comments (up to 5 levels).
- Mobile responsiveness.
- Progressive web app (allowing users to install the web app onto their devices).
- Searching for forums based on title, description, tags, and author username.
- A point system (Aura) that will increase when a user's thread or comment gets liked, incentivizing participation.

### Planned or In-Progress Features

- Content moderation using AI and community feedback.
- AI-generated descriptions for threads.
- Communities feature with each community having its own theme and color scheme.

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
<img width="1080" alt="data_model" src="https://github.com/user-attachments/assets/fc3edd3a-a598-4d0f-b739-e562bad081f0" />
  
*Figure: The ERD diagram for Uniconn's relational database.*
</div>

<br>

## Architecture

<div align="center">
<img width="601" alt="architecture" src="https://github.com/user-attachments/assets/8bdd8aa7-3c2e-40de-aff5-7bed5962d037" />
  
*Figure: The tech architecture for Uniconn.*
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

<br>

## Challenges

1. **Cross-Origin Auth Cookie**:  
   Different browsers handle these cookies differently. Initially, the cookie only worked on Firefox, and after some adjustments, it now works on Chrome as well. However, it still does not work on Safari.

2. **UI/UX Design**:  
   I am not very proficient in frontend design, so it took time to get the UI/UX (sort of?) right. I have settled on a decent design for now but might update it later.

3. **Code Quality**:  
   The current code looks like it was written by a competitive programmer (I am not one) `\s`. Over the next two weeks, I plan to refactor the codebase and follow SOLID principles to improve code quality.

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

4. Add a .env file. In the .env file, add the following:

   - DATABASE_URL
   - JWT_SECRET
   - REDIS_URL
   - REDIS_PWD
   - REDIS_USERNAME

5. Start the server. Run:

```console
go run .
```

6. Start a new terminal session.

7. Change your directory to the frontend. Run:

```console
cd gossip/frontend
```

8. Download the required modules. Run:

```console
npm install
```

9. Start the server. Run:

```console
npm run dev
```

10. The server will start at localhost:5173.

<br>

## Appreciation

Thanks to the CVWO Team for this awesome assignment. Although I had some previous experience, this assignment allowed me to fill in the gaps in my knowledge and push my boundaries.
