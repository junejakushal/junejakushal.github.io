# sv

Everything you need to build a Svelte project, powered by [`sv`](https://github.com/sveltejs/cli).

## Creating a project

If you're seeing this, you've probably already done this step. Congrats!

```sh
# create a new project
npx sv create my-app
```

To recreate this project with the same configuration:

```sh
# recreate this project
npx sv@0.12.5 create --template minimal --types ts --install npm .
```

## Developing

Once you've created a project and installed dependencies with `npm install` (or `pnpm install` or `yarn`), start a development server:

```sh
npm run dev

# or start the server and open the app in a new browser tab
npm run dev -- --open
```

## Building

To create a production version of your app:

```sh
npm run build
```

You can preview the production build with `npm run preview`.

> To deploy your app, you may need to install an [adapter](https://svelte.dev/docs/kit/adapters) for your target environment.

## Preview commands

This project uses `pnpm`.

Start the development server (live reload, draft posts shown):

```sh
pnpm run dev
```

Then open the printed URL (usually `http://localhost:5173`).

Build and preview the production site:

```sh
pnpm run build
pnpm run preview
```

Then open the printed URL (usually `http://localhost:4173`).

The MDX test page with the interactive digit canvas is at `/writings/writings-tab-test`.
