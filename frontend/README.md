# Running

Be sure that you've installed [docker](https://docs.docker.com/desktop/) and [docker-compose](https://docs.docker.com/compose/install/) or/and [node](https://nodejs.org/en).

## Using docker compose

```sh
$ docker-compose up --build
```

## Using npm

```sh
$ npm install
$ npm run dev
```

# Algorith & Approach

## Libraries

- [Vite](https://vitejs.dev/): A front-end build tooling that provides a quick ReactTs template.
- [Axios](https://axios-http.com/): An HTTP client to make HTTP requests.
- [styled-components](https://styled-components.com/): A library that facilitates creating React Components with styles.
- [Material Icons](https://mui.com/material-ui/material-icons/): Provide a variety of Icons.

## Architecture

The project architecture is divided in `components`, `pages`, `helpers` and `types`.

- [components](./src/components/): The components contains all the visual elements that can be found in the project. They are built using dependency injection so we can reuse them in different contexts not tying it to a specific usage, but, giving it a general usage. For example, the `Input` component, that's used for the search bar and the password form.
- [pages](./src/pages/): The pages are the combination of components.
- [helpers](./src/helpers/): Helpers are a piece of code that can be reused across the project.
- [types](./src/types/): Types are representations of the project's entities.

This way we can grow the project in a way that can be extensible and testable.

## Limitations

- Tests - No tests in the frontend.
- Pagination - No pagination built.
