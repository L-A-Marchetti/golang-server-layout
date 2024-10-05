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
    print_color "yellow" "Warning: You are on the $main_branch branch!"
    print_color "green" "To switch to a different branch, use: git checkout <branch_name>"
    print_color "green" "To create a new branch, use: git checkout -b <new_branch>"
    print_color "green" "Then, to set the upstream and push the new branch:"
    print_color "green" "git push -u origin <new_branch>"
    exit 1
fi

# Ask user if they want to update and merge
print_color "yellow" "You are currently on branch: $current_branch"
print_color "yellow" "Do you want to update $main_branch and merge it into your current branch?"
print_color "yellow" "This will bring your branch up to date with $main_branch, but may cause conflicts."
read -p "Update and merge? (y/n): " answer

if [[ $answer =~ ^[Yy]$ ]]; then
    # Update the main branch
    print_color "yellow" "Updating the $main_branch branch..."
    git checkout $main_branch
    git fetch origin
    git pull origin $main_branch

    # Return to the working branch and update it
    print_color "yellow" "Updating your working branch..."
    git checkout $current_branch
    git merge $main_branch

    print_color "green" "Your branch is now up to date with $main_branch."
    print_color "yellow" "If there were any merge conflicts, please resolve them now."
else
    print_color "yellow" "Skipping update and merge. Your branch may be behind $main_branch."
fi

print_color "green" "You're ready to start working on branch: $current_branch"
print_color "yellow" "Remember to commit and push regularly to save your work:"
print_color "green" "git add ."
print_color "green" "git commit -m \"Description of your changes\""
print_color "green" "git push origin $current_branch"