# ResIT Frontend

<a href="https://github.com/RealSnowKid/ResIT">
    <img src="https://raw.githubusercontent.com/RealSnowKid/ResIT/master/img/logo_horizontal.png" alt="Logo" width="170" height="55">
</a>

## Recommended IDE Setup

[VSCode](https://code.visualstudio.com/) + [Volar](https://marketplace.visualstudio.com/items?itemName=johnsoncodehk.volar) (and disable Vetur) + [TypeScript Vue Plugin (Volar)](https://marketplace.visualstudio.com/items?itemName=johnsoncodehk.vscode-typescript-vue-plugin).

## Type Support for `.vue` Imports in TS

TypeScript cannot handle type information for `.vue` imports by default, so we replace the `tsc` CLI with `vue-tsc` for type checking. In editors, we need [TypeScript Vue Plugin (Volar)](https://marketplace.visualstudio.com/items?itemName=johnsoncodehk.vscode-typescript-vue-plugin) to make the TypeScript language service aware of `.vue` types.

## Build with

 - Vue 3
 - Vite
 - Vue Router
 - Vue Axios
 - Vuex
 - Tailwind

## How to run

1. Make sure to have NodeJS and NPM installed
2. First install the node_modules with `npm install`
3. To run the site with hot-reload `npm run dev`
4. To build the site for production `npm run build`

## Folder Structure

```c++
├───public          // Basic HTML files
├───src             // Source code
    ├───assets      // Logos, css
    ├───components  // Components that can be used by views
    ├───router      // Http router
    └───views       // Full pages that use the components

```

