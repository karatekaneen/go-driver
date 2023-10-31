# Change Log

## [master](https://github.com/arangodb/go-driver/tree/master) (N/A)

## [2.0.3](https://github.com/arangodb/go-driver/tree/v2.0.3) (2023-10-31)
- Add optional status code checks. Consistent return of response
- JavaScript Transactions API
- Async Client
- Fix connection.NewRequestWithEndpoint()
- Add support for MaglevHashEndpoints
- Add basic support for Views and Analyzers
- Add ServerMode/SetServerMode/ServerID
- Add collection Truncate, Count, Properties, SetProperties
- Add and re-organize missing collection properties fields
- Rename CreateCollectionOptions to CreateCollectionProperties
- Add support for missing query options (create documents, remove collection, remove view)
- Adjust CursorStats and JournalSize types
- Improve returning old doc handling in CollectionDocumentDelete
- Agency: Supply ClientID with agency transactions
- Automate release process