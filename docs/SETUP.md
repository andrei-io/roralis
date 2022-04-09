# RO Setup pentru mobile

## Setup enviroment

Trebuie instalat:

- [node varianta LTS](https://nodejs.org/en/)
- yarn:
  - se deschide cmd dupa ce sa instalat node
  - `npm install --global yarn`
- [expo](https://docs.expo.dev/get-started/installation/)
  - cli : `npm install --global expo-cli`
  - [appplicatie mobila](https://play.google.com/store/apps/details?id=host.exp.exponents)
- Recomandat pe windows: git bash

## Setup editor

Extensii recomandate

- adpyke.vscode-sql-formatter
- bungcip.better-toml
- dbaeumer.vscode-eslint
- EditorConfig.EditorConfig
- esbenp.prettier-vscode
- golang.go
- Gruntfuggly.todo-tree
- mads-hartmann.bash-ide-vscode
- ms-python.python
- ms-vsliveshare.vsliveshare
- msjsdiag.vscode-react-native
- PKief.material-icon-theme
- VisualStudioExptTeam.vscodeintellicode
- wix.vscode-import-cost

Instalare rapida:

```bash
code --install-extension adpyke.vscode-sql-formatter && code --install-extension bungcip.better-toml && code --install-extension dbaeumer.vscode-eslint && code --install-extension EditorConfig.EditorConfig && code --install-extension esbenp.prettier-vscode && code --install-extension golang.go && code --install-extension Gruntfuggly.todo-tree && code --install-extension mads-hartmann.bash-ide-vscode && code --install-extension ms-python.python && code --install-extension ms-vsliveshare.vsliveshare && code --install-extension msjsdiag.vscode-react-native && code --install-extension PKief.material-icon-theme && code --install-extension VisualStudioExptTeam.vscodeintellicode && code --install-extension wix.vscode-import-cost
```

Setari necesare:

Ctrl-shift-p > Preferences: Open Settings (JSON)

In fisier se pun urmatoarele:

```json
{
  "[javascript]": {
    "editor.defaultFormatter": "esbenp.prettier-vscode",
    "editor.codeActionsOnSave": {
      "source.organizeImports": true
    }
  },
  "[typescript]": {
    "editor.defaultFormatter": "esbenp.prettier-vscode",
    "editor.codeActionsOnSave": {
      "source.organizeImports": true
    }
  },
  "[typescriptreact]": {
    "editor.defaultFormatter": "esbenp.prettier-vscode",
    "editor.codeActionsOnSave": {
      "source.organizeImports": true
    }
  },
  "[json]": {
    "editor.defaultFormatter": "esbenp.prettier-vscode"
  },
  "javascript.preferences.quoteStyle": "single",
  "typescript.preferences.quoteStyle": "single",
  "editor.insertSpaces": false,
  "[go]": {
    "editor.insertSpaces": false,
    "editor.formatOnSave": true,
    "editor.codeActionsOnSave": {
      "source.organizeImports": true
    },
    "editor.defaultFormatter": "golang.go",
    "editor.tabSize": 4
  },
  "go.lintTool": "golangci-lint",
  "go.lintFlags": ["--fast"],
  "typescript.updateImportsOnFileMove.enabled": "prompt",
  "prettier.printWidth": 100,
  "prettier.trailingComma": "all",
  "prettier.singleQuote": true,
  "todo-tree.general.tags": ["BUG", "HACK", "FIXME", "TODO", "XXX", "REFACTOR"],
  "todo-tree.regex.regex": "(//|#|<!--|;|/\\*|^|^\\s*(-|\\d+.))\\s*($TAGS)",
  "[jsonc]": {
    "editor.defaultFormatter": "esbenp.prettier-vscode"
  },
  "files.watcherExclude": {
    "**/.git/objects/**": true,
    "**/.git/subtree-cache/**": true,
    "**/node_modules/*/**": true,
    "**/.hg/store/**": true,
    "**/docker_data/**": true
  },
  "explorer.compactFolders": false,
  "editor.formatOnSave": true,
  "javascript.preferences.useAliasesForRenames": false,
  "typescript.preferences.useAliasesForRenames": false,
  "files.exclude": {
    "**/node_modules": true
  }
}
```

## Setup proiect

Aplicatia mobila se afla in folderul **mobile**

Trebuie instalate dependentele proiectului: `yarn`

Exista mai multe optiuni optiuni

Rulat din terminal: `yarn start`

Rulat din VS Code(recomandat):

- Ctrl-shift-p > Task: Run Task > yarn: start - mobile
- va aparea un terminal in vs code, unde vor fi printate instructiuni pe ecran
