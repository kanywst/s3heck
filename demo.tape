# s3heck demo tape
# Run with: vhs demo.tape

Output demo.gif

Set Shell "bash"
Set FontSize 18
Set Width 1200
Set Height 600
Set Padding 40
Set FontFamily "Menlo"
Set Theme "Catppuccin Mocha"
Set WindowBar Colorful

# Setup
Hide
Type "go build -o s3heck ."
Enter
Type "alias s3='./s3heck'"
Enter
Type "clear"
Enter
Show

# 1. Introduction
Type "s3 --help"
Sleep 500ms
Enter
Sleep 2s

# 2. Remote Diagnostics (The "Wow" Factor)
Type "clear"
Enter
Type "# 1. Instantly diagnose TLS chains & validity"
Sleep 1s
Enter
Type "s3 check --remote google.com"
Sleep 500ms
Enter
Sleep 3s

# 3. Scaffolding (Solving the "How do I start?" problem)
Type "clear"
Enter
Type "# 2. Generate production-ready configuration"
Sleep 1s
Enter
Type "s3 scaffold autocert > deployment.yaml"
Sleep 500ms
Enter
Sleep 1s
Type "cat deployment.yaml"
Enter
Sleep 3s

# 4. Learning (Integrated Docs)
Type "clear"
Enter
Type "# 3. Access ecosystem docs without leaving the terminal"
Sleep 1s
Enter
Type "s3 learn kms | head -n 15"
Enter
Sleep 3s

# Outro
Type "clear"
Enter
Type "# s3: The Modern PKI Orchestrator"
Sleep 1s
Enter
Sleep 2s
