#!/bin/bash

# Function to print colored messages
print_color() {
    case $1 in
        "red") COLOR='\033[0;31m' ;;
        "green") COLOR='\033[0;32m' ;;
        "yellow") COLOR='\033[1;33m' ;;
        *) COLOR='\033[0m' ;;
    esac
    NC='\033[0m' # No Color
    echo -e "${COLOR}$2${NC}"
}

# Determine the main branch name (main or master)
if git show-ref --quiet refs/heads/main; then
    main_branch="main"
elif git show-ref --quiet refs/heads/master; then
    main_branch="master"
else
    print_color "red" "Error: Neither 'main' nor 'master' branch found."
    exit 1
fi

# Check current branch
current_branch=$(git rev-parse --abbrev-ref HEAD)
if [ "$current_branch" = "$main_branch" ]; then
    print_color "red" "Error: You are on the $main_branch branch!"
    print_color "yellow" "Please switch to your feature branch before stopping work."
    print_color "green" "Use: git checkout <your_feature_branch>"
    exit 1
fi

# Check for untracked files
if [ -n "$(git ls-files --others --exclude-standard)" ]; then
    print_color "yellow" "Warning: You have untracked files."
    print_color "green" "To see these files, use: git status"
    print_color "green" "To add all new files, use: git add ."
    print_color "green" "To add specific files, use: git add <file1> <file2> ..."
    print_color "yellow" "Please add these files if they should be part of your commit."
fi

# Check for unstaged changes
if ! git diff --quiet; then
    print_color "yellow" "Warning: You have unstaged changes."
    print_color "green" "To see these changes, use: git status"
    print_color "green" "To stage all changes, use: git add ."
    print_color "green" "To stage specific files, use: git add <file1> <file2> ..."
    print_color "yellow" "Please stage these changes if they should be part of your commit."
fi

# Check for uncommitted changes
if ! git diff --staged --quiet; then
    print_color "yellow" "Warning: You have staged but uncommitted changes."
    print_color "green" "Please commit your changes before creating a pull request:"
    print_color "green" "git commit -m \"Description of your changes\""
fi

# Check if all commits have been pushed
if [ "$(git rev-parse HEAD)" != "$(git rev-parse @{u})" ]; then
    print_color "yellow" "Warning: You have local commits that haven't been pushed."
    print_color "green" "Please push your commits before creating a pull request:"
    print_color "green" "git push origin $current_branch"
fi

# Final check to see if everything is ready
if [ -z "$(git status --porcelain)" ] && [ "$(git rev-parse HEAD)" == "$(git rev-parse @{u})" ]; then
    print_color "green" "Your branch '$current_branch' is ready for review!"
    print_color "yellow" "To create a pull request on Gitea:"
    print_color "green" "1. Go to your Gitea repository in your web browser"
    print_color "green" "2. Click on 'Pull Requests' in the menu"
    print_color "green" "3. Click on 'New Pull Request'"
    print_color "green" "4. Select your branch '$current_branch' as 'source' and '$main_branch' as 'target'"
    print_color "green" "5. Add a descriptive title and comments about your changes"
    print_color "green" "6. Click on 'Create Pull Request'"

    print_color "yellow" "Don't forget to:"
    print_color "green" "- Clearly describe the features you've added or modified"
    print_color "green" "- Mention any issues resolved, if applicable"
    print_color "green" "- Request a review from your teammates or project manager"
    print_color "green" "- Respond to comments and make necessary changes"

    print_color "yellow" "Great job! Don't hesitate to ask for help if you have any questions."
else
    print_color "yellow" "Please address the warnings above before creating a pull request."
fi