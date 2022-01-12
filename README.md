# android_env_setup_1337

add these lines to .zshenv (if not exist create it)
```zsh
# custom paths
export DEV_TOOLS=/Users/$(whoami)/goinfre/devtools
export APPS=/Users/$(whoami)/goinfre/apps

export ANDROID_SDK_ROOT=$DEV_TOOLS/Android/sdk
export GRADLE_USER_HOME=$DEV_TOOLS/gradle
export GRADLE_OPTS=-Dgradle.user.home=$GRADLE_USER_HOME

# java home
export JAVA_17_HOME=$DEV_TOOLS/jdk/Contents/Home
export JAVA_HOME=$JAVA_17_HOME
```

.zshrc
```zsh
export PATH=$PATH:$DEV_TOOLS/Android/sdk/tools/bin
export PATH=$PATH:$DEV_TOOLS/Android/sdk/emulator
export PATH="$PATH:$DEV_TOOLS/flutter/bin"
export PATH="$PATH:$DEV_TOOLS/gradle/bin"
export PATH="$PATH:/Applications/Visual Studio Code.app/Contents/Resources/app/bin"

# java
export PATH=$PATH:$JAVA_HOME/bin
```

android sdk path:
```
/Users/'your user name'/goinfre/devtools/Android/sdk
```
download command line tools + android apis
