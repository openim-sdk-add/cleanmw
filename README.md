Use this SDK to add instant messaging capabilities to your app. By connecting to a self-hosted OpenIM server, you can quickly integrate instant messaging capabilities into your app with just a few lines of code.

The underlying SDK core is implemented in OpenIM SDK Core. Using gomobile, it can be compiled into an XCFramework for iOS integration. iOS interacts with the OpenIM SDK Core through JSON, and the SDK exposes a re-encapsulated API for easy usage. In terms of data storage, iOS utilizes the SQLite layer provided internally by the OpenIM SDK Core.
