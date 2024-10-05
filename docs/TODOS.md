# Forum Project Task List

## Week 1

### Day 1-2: Project Setup and Database

1. Set up project structure
   - Create directories for Go backend and HTML frontend
   - Initialize Git repository
2. Initialize Go module
   - Run `go mod init <module_name>`
3. Create `main.go` file
   - Set up basic HTTP server
4. Implement database structure based on DBML
   - Create SQL scripts for table creation (users, posts, comments, categories, post_categories, likes_dislikes, images)
5. Implement database connection in Go
   - Choose and install PostgreSQL driver (e.g., `pq`)
   - Write connection logic in `main.go`
6. Implement basic CRUD functions for each table
7. Set up database encryption
   - Research and implement encryption method
8. Implement password system for database access

cf [Database table creation and setup](CREATE_TABLES.md)

cf [Creating CRUD functions](CRUD.md)

cf [Data encryption, hashed passwords & sqlite security](ENCRYPTION.md)

### Day 3-4: Authentication System

9. Create HTML login page
   - Design form with email and password fields
10. Implement server-side login function
    - Verify email and password against the database
    - Return appropriate error messages for failed attempts
11. Create HTML registration page
    - Design form with email, username, and password fields
12. Implement server-side registration function
    - Check if email is already taken
    - Encrypt password before storing
13. Integrate Google authentication
    - Configure Google API
    - Implement server-side OAuth flow
    - Create Google login button on client-side
14. Integrate GitHub authentication
    - Configure GitHub API
    - Implement server-side OAuth flow
    - Create GitHub login button on client-side
15. Implement session management with UUID cookies
    - Generate UUIDs for user sessions
    - Store session state securely on the server
16. Create different user types (Guests, Users, Moderators, Administrators)
    - Define user roles and permissions in the database

cf [Using go templates](TEMPLATES.md)

cf [oAuth for login with google & github](OAUTH.md)

cf [session cookies using UUID](UUID.md)

### Day 5-6: User Management

17. Create user profile page
    - Design layout to display user information
    - Implement server-side logic to fetch user data
18. Implement profile editing functionality
    - Create form for users to update their information
    - Implement server-side logic to update user data
19. Create admin interface for user role management
    - Design admin dashboard layout
    - Implement functionality to view all users
    - Add ability to change user roles (User, Moderator, Administrator)
20. Implement moderator role request functionality
    - Create interface for users to submit moderator requests
    - Implement server-side logic to store and process requests
21. Create interface for admins to manage moderator role requests
    - Design page to display pending moderator requests
    - Implement approve/deny functionality for requests
22. Implement user activity tracking
    - Create database structure to store user actions
    - Implement logic to record user posts, comments, likes/dislikes
23. Design and implement user permissions system
    - Define actions allowed for each user type
    - Implement server-side checks for user permissions

### Day 7: Post Management

24. Create HTML post creation page
    - Design form with title, content, and category selection
25. Implement server-side post creation function
    - Verify post is not empty
    - Allow selection of multiple categories
26. Implement image upload functionality
    - Set up file handling for image uploads
    - Limit size to 20MB and validate file types (.jpg, .png, .gif)
27. Create HTML post display page
28. Implement server-side post retrieval function
29. Create post editing functionality
30. Implement post deletion
31. Set up post validation system for new users

## Week 2

### Day 8-9: Comment System and Like/Dislike

32. Create HTML interface for comments
    - Design comment section layout under each post
33. Implement server-side comment addition function
    - Ensure only registered users can comment
34. Create comment editing functionality
    - Allow users to edit their own comments
35. Implement comment deletion
    - For users (their own comments) and moderators/admins
36. Create like/dislike buttons for posts and comments
    - Design visually distinguishable buttons
37. Implement server-side like/dislike logic
    - Prevent simultaneous liking and disliking
    - Update database to reflect user actions
38. Update like count display in real-time
    - Implement AJAX for dynamic like/dislike updates

### Day 10: Filters, Sorting, and Categories

39. Implement category-based filter
    - Create UI for category selection
    - Implement server-side filtering logic
40. Create sorting functionality by creation date
    - Add sorting options to the UI
    - Implement server-side sorting logic
41. Implement sorting by number of likes
    - Extend sorting options in the UI
    - Add server-side logic for like-based sorting
42. Create interface for category management (admin only)
    - Design category management page
    - Implement category creation functionality
43. Implement category deletion
    - Add delete option for each category
    - Implement server-side deletion logic

### Day 11: Activity and Notifications

44. Create user activity page
    - Design layout to display user's posts, comments, and likes
    - Implement server-side logic to fetch activity data
45. Implement notification system for likes, dislikes, and comments
    - Create database structure for notifications
    - Implement logic to generate notifications on user actions
46. Create interface to display notifications
    - Design notification panel/page
    - Implement real-time notification updates

### Day 12: Moderation

47. Create reporting interface for posts and comments
    - Add report buttons to posts and comments
    - Implement server-side logic to store reports
48. Implement server-side reporting logic
    - Create database structure for reports
    - Implement logic to process and store report categories
49. Create moderator interface for managing reports
    - Design dashboard for moderators to view and handle reports
    - Implement functionality to resolve reports
50. Implement post and comment deletion functionality for moderators and admins
    - Add delete options in moderator/admin interfaces
    - Implement server-side deletion logic with proper permissions

### Day 13: Security

51. Configure HTTPS using autocert package or OpenSSL
    - Set up SSL certificate generation
    - Configure server to use HTTPS
52. Implement rate limiting system
    - Research and choose a rate limiting library
    - Implement rate limiting for API endpoints
53. Configure Go TLS structure
    - Set up proper TLS configuration in Go
54. Reduce server timeout
    - Configure appropriate timeout settings
55. Create configuration system for certificate information
    - Implement .env or config file system
56. Create custom pages for 404, 400, and 500 errors
    - Design error pages
    - Implement server-side logic to serve custom error pages
57. Implement client session cookies with unique identifiers
    - Ensure secure cookie handling
    - Implement server-side session validation

### Day 14: Final Touches and Docker

58. Optimize database queries
    - Review and optimize complex queries
    - Implement database indexing where necessary
59. Conduct security testing
    - Perform penetration testing
    - Address any security vulnerabilities found
60. Write API documentation
    - Document all API endpoints and their usage
61. Create user manual
    - Write comprehensive guide for using the forum
62. Set up Docker configuration
    - Create Dockerfile for the application
    - Set up docker-compose for easy deployment
63. Perform final testing and bug fixes
    - Conduct thorough system testing
    - Fix any remaining bugs or issues
64. Deploy application
    - Prepare production environment
    - Deploy the application to the production server
