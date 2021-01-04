srv
================================================================================

The Simplest Static Fileserver

`srv` is the simplest fileserver I could imagine.

Usage
--------------------------------------------------------------------------------

Serve the current working directory on `:8000`:
```bash
srv
```

Serve a directory on `:8000`:
```bash
srv -dir 'files'
```

Serve a directory on a specified port:
```bash
srv -dir 'files' -addr ':8001'
```
