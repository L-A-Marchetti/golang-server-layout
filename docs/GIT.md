## Git Flow Process on Gitea

1. **Review Open Issues**
   - Go to the Issues section in Gitea
   - Look for unassigned issues

2. **Assign Yourself to an Issue**
   - Choose an issue and assign it to yourself
   - Note the issue number (e.g., #42)

3. **Create a Branch**
   - Use Gitea to create a new branch from master
   - Name the branch referencing the issue number (e.g., "feature-42-add-login")

4. **Ensure You're on Master Branch**
   ```bash
   git branch
   ```
   - If not on master:
     ```bash
     git checkout master
     ```

5. **Update Master Branch**
   ```bash
   git fetch
   git pull
   ```

6. **Switch to New Branch Locally**
   ```bash
   git checkout <branch-name>
   ```

7. **Regular Commits**
   - Make changes and commit frequently:
     ```bash
     git add -A
     git commit -m "#42 Add login form"
     git push
     ```
   - Always reference the issue number with a hashtag (#) in commit messages

8. **Complete Feature**
   - When the feature is complete, don't merge automatically

9. **Create Pull Request**
   - Go back to Gitea
   - Navigate to your branch
   - Click on "New Pull Request"
   - Name the pull request, referencing the issue number (e.g., "Feature #42: Add login functionality")

10. **Code Review**
    - Wait for a team member to review your code
    - Address any feedback or changes requested

11. **Merge**
    - Once approved, the pull request can be merged into the master branch
