# Summary for the forum project

cf [GIT WORKFLOW PROCESS #VERY IMPORTANT#](docs/GIT.md)

cf [Complete todo list for more details on tasks and sub-tasks](docs/TODOS.md)

### Database

cf [online DBML](https://dbdiagram.io/d/670128f4fb079c7ebd6e0d63)

cf [DBML.pdf](DBML.pdf)

### Authentification

- Login:
  - Must ask for email (return an error if the email is already taken)
  - Must ask for username
  - Must ask for password (password need to be encrypted)
- Register:
  - User should be able to register and login using Google and GitHub
- Cookies (UUID)

- User types:
  - We should implement at least 4 types of users:
    - Guests: Can't interact with the forum, only see the content
    - Users: Able to create, comment, like/dislike posts, but their posts are not instantly visible and should be validated by a moderator
    - Moderators: Should be able to delete or report posts to the admins. The report should be attached with these categories (irrelevant, obscene, illegal or insulting). Users can become moderators by requesting it, admins should be able to accept or not that role
    - Administrators: Should be able to change the role of any users. Receive reports from moderators and be able to respond. Be able to respond to moderators role requests from users. Administrators should also be able to delete posts and comments. But also be able to create and delete categories

### Communication

- Posts:
  - Only registered users can post
  - You can't post an empty post.
  - User need to select a category where to create the post, several categories can be selected.
  - Posts should be visible for all users
  - User should be able to upload an image into the post with at least .jpg, .png & .gif (the max size of the img should be 20mb. if it's superior then return an error)
  - User should be able to edit/remove their own posts, likes/dislikes and comments
- Comments:
  - Only registered users can comment
  - Comments should be visible for all users
- Likes/Dislikes:
  - You can't like and dislike something at the same time.
  - Posts and comments can be liked and disliked
  - Only registered users can like/dislike
  - The number of likes should be visible for all users
- Filters:
  - Categories
  - Last created
  - Most liked
- Activity:
  - Users should be notified when their posts are liked/disliked or commented
  - Create an activity page where the user can monitor their posts, their likes/dislikes, where and what the user has been commenting

### Security

- Implement a Hypertext Transfer Protocol Secure (HTTPS) (generate an SSL certificate). (cf. autocert package go, openssl manual)
- Rate limiting
- Database should be encrypted and have a password to access it.
- Client session cookies should be unique, the session state is stored on the server and the session should present a unique identifier. (cf. UUID, bcrypt)
- Handle 404, 400 & 500 errors.
- Go TLS structure configuration.
- Server Timeout reduced.
- A way to configure the certificates information, either via .env, config files or another method

### Docker
