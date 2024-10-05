# Git Helper Scripts

This repository contains two bash scripts designed to help beginners with their Git workflow. These scripts will guide you through starting your work session and preparing your changes for review.

## Prerequisites

- Git installed on your system
- Basic understanding of Git concepts (branches, commits, push/pull)

## Usage

### Starting Your Work Session

Run the `pre-work-git.sh` script when you're ready to begin working:

```
./pre-work-git.sh
```

This script will:
1. Check which branch you're on and alert you if you're in the main one.
2. Offer to update your branch with the latest changes from the main branch
3. Guide you through creating a new branch if needed
4. If you were in the main and created a new branch following the guide please relaunch again the `pre-work-git.sh` script to be sure your branch is updated with the main one.

### Ending Your Work Session

Run the `post-work-git.sh` script when you've finished your work and want to submit it for review:

```
./post-work-git.sh
```

This script will:
1. Check your current branch
3. Check for any untracked, unstaged, or uncommitted changes
4. Ensure all your commits are pushed
5. Guide you through creating a pull request if everything is in order

## Tips for Beginners

1. Always run `pre-work-git.sh` before you begin coding. This ensures you're working on the right branch with the latest updates. Also IF YOU WERE ON THE MAIN BRANCH PLEASE RELAUNCH THE SCRIPT AFTER YOU SWITCH TO YOUR OWN BRANCH TO BE SURE YOUR BRANCH IS UPDATED WITH THE MAIN.

2. Commit your changes regularly. A good practice is to commit each logical chunk of work.

3. Push your commits to the remote repository frequently to back up your work.

4. Before ending your session, always run `post-work-git.sh`. This script will catch any forgotten steps and guide you through the process of submitting your work for review.

5. If you encounter any errors or warnings while running these scripts, read the messages carefully. They often provide instructions on how to resolve the issue.

6. Don't hesitate to ask for help from your team members or project manager if you're unsure about anything.

Remember, these scripts are helpers, but understanding Git is still important. Take some time to learn about Git concepts and commands to become more proficient in your development workflow.

Happy coding!