# Jekyll Site Setup - Deployment Instructions

This repository has been configured with Jekyll Static Site Generator using the minima theme for GitHub Pages deployment.

## What's Been Set Up

### Core Files
- **Gemfile**: Specifies Jekyll and GitHub Pages dependencies
- **_config.yml**: Jekyll configuration with minima theme
- **index.md**: Home page
- **about.md**: About page
- **.gitignore**: Excludes build artifacts and vendor files

### Content
- **_posts/2025-11-07-welcome-to-jekyll.md**: Sample blog post

### GitHub Actions
- **.github/workflows/jekyll.yml**: Automated build and deployment workflow

## GitHub Pages Setup Instructions

To deploy this site to GitHub Pages, you need to:

1. **Enable GitHub Pages** in your repository:
   - Go to repository Settings → Pages
   - Under "Build and deployment":
     - Source: Select "GitHub Actions"
   - Save the settings

2. **Merge this PR** to the main branch:
   - Once merged, the GitHub Actions workflow will automatically trigger
   - The site will be built and deployed to GitHub Pages
   - The deployment typically takes 1-2 minutes

3. **Access your site**:
   - Your site will be available at: `https://junejakushal.github.io/`

## Local Development

To run the site locally:

```bash
# Install dependencies (first time only)
bundle install

# Serve the site locally
bundle exec jekyll serve

# The site will be available at http://localhost:4000
```

## Customization

You can customize the site by editing:

- **_config.yml**: Change site title, description, social media links, etc.
- **index.md**: Modify the home page content
- **about.md**: Update the about page
- Add new posts in **_posts/** following the naming convention: `YYYY-MM-DD-title.md`

## Resources

- [Jekyll Documentation](https://jekyllrb.com/docs/)
- [Minima Theme Documentation](https://github.com/jekyll/minima)
- [GitHub Pages Documentation](https://docs.github.com/en/pages)
