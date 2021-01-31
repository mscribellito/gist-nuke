# gist-nuke

Remove all gists from a GitHub account.

**Caution!** *gist-nuke* is a destructive tool. By default *gist-nuke* only lists gists. You need to add `--no-dry-run` to actually remove gists.

## Usage

List all gists and exit:

```
> gist-nuke --token <github-token>
Fancy Gist - '8e94bba2b9c6a4285e31438804c562d7' - would remove
```

Tries to delete all gists:

```
gist-nuke --token <github-token> --no-dry-run
Fancy Gist - '8e94bba2b9c6a4285e31438804c562d7' - finished

Removal: 1 removed, 0 failed
```

### GitHub Personal Access Token

*gist-nuke* requires a GitHub [Personal Access Token](https://github.com/settings/tokens) with delete_repo and gist [scopes](https://docs.github.com/en/developers/apps/scopes-for-oauth-apps#available-scopes).
