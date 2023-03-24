# Go API client for whisparr

Radarr API docs

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

[comment]: # (x-release-please-start-version)
- Package version: 0.2.0

[comment]: # (x-release-please-end)
- API version: 3.0.0

- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import whisparr "github.com/devopsarr/whisparr-go"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), whisparr.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), whisparr.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), whisparr.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), whisparr.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost:7878*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AlternativeTitleApi* | [**GetAlttitleById**](whisparr/docs/AlternativeTitleApi.md#getalttitlebyid) | **Get** /api/v3/alttitle/{id} | 
*AlternativeTitleApi* | [**ListAlttitle**](whisparr/docs/AlternativeTitleApi.md#listalttitle) | **Get** /api/v3/alttitle | 
*ApiInfoApi* | [**GetApi**](whisparr/docs/ApiInfoApi.md#getapi) | **Get** /api | 
*AuthenticationApi* | [**CreateLogin**](whisparr/docs/AuthenticationApi.md#createlogin) | **Post** /login | 
*AuthenticationApi* | [**GetLogout**](whisparr/docs/AuthenticationApi.md#getlogout) | **Get** /logout | 
*BackupApi* | [**CreateSystemBackupRestoreById**](whisparr/docs/BackupApi.md#createsystembackuprestorebyid) | **Post** /api/v3/system/backup/restore/{id} | 
*BackupApi* | [**CreateSystemBackupRestoreUpload**](whisparr/docs/BackupApi.md#createsystembackuprestoreupload) | **Post** /api/v3/system/backup/restore/upload | 
*BackupApi* | [**DeleteSystemBackup**](whisparr/docs/BackupApi.md#deletesystembackup) | **Delete** /api/v3/system/backup/{id} | 
*BackupApi* | [**ListSystemBackup**](whisparr/docs/BackupApi.md#listsystembackup) | **Get** /api/v3/system/backup | 
*BlocklistApi* | [**DeleteBlocklist**](whisparr/docs/BlocklistApi.md#deleteblocklist) | **Delete** /api/v3/blocklist/{id} | 
*BlocklistApi* | [**DeleteBlocklistBulk**](whisparr/docs/BlocklistApi.md#deleteblocklistbulk) | **Delete** /api/v3/blocklist/bulk | 
*BlocklistApi* | [**GetBlocklist**](whisparr/docs/BlocklistApi.md#getblocklist) | **Get** /api/v3/blocklist | 
*BlocklistApi* | [**ListBlocklistMovie**](whisparr/docs/BlocklistApi.md#listblocklistmovie) | **Get** /api/v3/blocklist/movie | 
*CalendarApi* | [**GetCalendarById**](whisparr/docs/CalendarApi.md#getcalendarbyid) | **Get** /api/v3/calendar/{id} | 
*CalendarApi* | [**ListCalendar**](whisparr/docs/CalendarApi.md#listcalendar) | **Get** /api/v3/calendar | 
*CalendarFeedApi* | [**GetFeedV3CalendarRadarrIcs**](whisparr/docs/CalendarFeedApi.md#getfeedv3calendarradarrics) | **Get** /feed/v3/calendar/radarr.ics | 
*CollectionApi* | [**GetCollectionById**](whisparr/docs/CollectionApi.md#getcollectionbyid) | **Get** /api/v3/collection/{id} | 
*CollectionApi* | [**ListCollection**](whisparr/docs/CollectionApi.md#listcollection) | **Get** /api/v3/collection | 
*CollectionApi* | [**PutCollection**](whisparr/docs/CollectionApi.md#putcollection) | **Put** /api/v3/collection | 
*CollectionApi* | [**UpdateCollection**](whisparr/docs/CollectionApi.md#updatecollection) | **Put** /api/v3/collection/{id} | 
*CommandApi* | [**CreateCommand**](whisparr/docs/CommandApi.md#createcommand) | **Post** /api/v3/command | 
*CommandApi* | [**DeleteCommand**](whisparr/docs/CommandApi.md#deletecommand) | **Delete** /api/v3/command/{id} | 
*CommandApi* | [**GetCommandById**](whisparr/docs/CommandApi.md#getcommandbyid) | **Get** /api/v3/command/{id} | 
*CommandApi* | [**ListCommand**](whisparr/docs/CommandApi.md#listcommand) | **Get** /api/v3/command | 
*CreditApi* | [**GetCreditById**](whisparr/docs/CreditApi.md#getcreditbyid) | **Get** /api/v3/credit/{id} | 
*CreditApi* | [**ListCredit**](whisparr/docs/CreditApi.md#listcredit) | **Get** /api/v3/credit | 
*CustomFilterApi* | [**CreateCustomFilter**](whisparr/docs/CustomFilterApi.md#createcustomfilter) | **Post** /api/v3/customfilter | 
*CustomFilterApi* | [**DeleteCustomFilter**](whisparr/docs/CustomFilterApi.md#deletecustomfilter) | **Delete** /api/v3/customfilter/{id} | 
*CustomFilterApi* | [**GetCustomFilterById**](whisparr/docs/CustomFilterApi.md#getcustomfilterbyid) | **Get** /api/v3/customfilter/{id} | 
*CustomFilterApi* | [**ListCustomFilter**](whisparr/docs/CustomFilterApi.md#listcustomfilter) | **Get** /api/v3/customfilter | 
*CustomFilterApi* | [**UpdateCustomFilter**](whisparr/docs/CustomFilterApi.md#updatecustomfilter) | **Put** /api/v3/customfilter/{id} | 
*CustomFormatApi* | [**CreateCustomFormat**](whisparr/docs/CustomFormatApi.md#createcustomformat) | **Post** /api/v3/customformat | 
*CustomFormatApi* | [**DeleteCustomFormat**](whisparr/docs/CustomFormatApi.md#deletecustomformat) | **Delete** /api/v3/customformat/{id} | 
*CustomFormatApi* | [**GetCustomFormatById**](whisparr/docs/CustomFormatApi.md#getcustomformatbyid) | **Get** /api/v3/customformat/{id} | 
*CustomFormatApi* | [**GetCustomFormatSchema**](whisparr/docs/CustomFormatApi.md#getcustomformatschema) | **Get** /api/v3/customformat/schema | 
*CustomFormatApi* | [**ListCustomFormat**](whisparr/docs/CustomFormatApi.md#listcustomformat) | **Get** /api/v3/customformat | 
*CustomFormatApi* | [**UpdateCustomFormat**](whisparr/docs/CustomFormatApi.md#updatecustomformat) | **Put** /api/v3/customformat/{id} | 
*DelayProfileApi* | [**CreateDelayProfile**](whisparr/docs/DelayProfileApi.md#createdelayprofile) | **Post** /api/v3/delayprofile | 
*DelayProfileApi* | [**DeleteDelayProfile**](whisparr/docs/DelayProfileApi.md#deletedelayprofile) | **Delete** /api/v3/delayprofile/{id} | 
*DelayProfileApi* | [**GetDelayProfileById**](whisparr/docs/DelayProfileApi.md#getdelayprofilebyid) | **Get** /api/v3/delayprofile/{id} | 
*DelayProfileApi* | [**ListDelayProfile**](whisparr/docs/DelayProfileApi.md#listdelayprofile) | **Get** /api/v3/delayprofile | 
*DelayProfileApi* | [**UpdateDelayProfile**](whisparr/docs/DelayProfileApi.md#updatedelayprofile) | **Put** /api/v3/delayprofile/{id} | 
*DiskSpaceApi* | [**ListDiskSpace**](whisparr/docs/DiskSpaceApi.md#listdiskspace) | **Get** /api/v3/diskspace | 
*DownloadClientApi* | [**CreateDownloadClient**](whisparr/docs/DownloadClientApi.md#createdownloadclient) | **Post** /api/v3/downloadclient | 
*DownloadClientApi* | [**CreateDownloadClientActionByName**](whisparr/docs/DownloadClientApi.md#createdownloadclientactionbyname) | **Post** /api/v3/downloadclient/action/{name} | 
*DownloadClientApi* | [**DeleteDownloadClient**](whisparr/docs/DownloadClientApi.md#deletedownloadclient) | **Delete** /api/v3/downloadclient/{id} | 
*DownloadClientApi* | [**GetDownloadClientById**](whisparr/docs/DownloadClientApi.md#getdownloadclientbyid) | **Get** /api/v3/downloadclient/{id} | 
*DownloadClientApi* | [**ListDownloadClient**](whisparr/docs/DownloadClientApi.md#listdownloadclient) | **Get** /api/v3/downloadclient | 
*DownloadClientApi* | [**ListDownloadClientSchema**](whisparr/docs/DownloadClientApi.md#listdownloadclientschema) | **Get** /api/v3/downloadclient/schema | 
*DownloadClientApi* | [**TestDownloadClient**](whisparr/docs/DownloadClientApi.md#testdownloadclient) | **Post** /api/v3/downloadclient/test | 
*DownloadClientApi* | [**TestallDownloadClient**](whisparr/docs/DownloadClientApi.md#testalldownloadclient) | **Post** /api/v3/downloadclient/testall | 
*DownloadClientApi* | [**UpdateDownloadClient**](whisparr/docs/DownloadClientApi.md#updatedownloadclient) | **Put** /api/v3/downloadclient/{id} | 
*DownloadClientConfigApi* | [**GetDownloadClientConfig**](whisparr/docs/DownloadClientConfigApi.md#getdownloadclientconfig) | **Get** /api/v3/config/downloadclient | 
*DownloadClientConfigApi* | [**GetDownloadClientConfigById**](whisparr/docs/DownloadClientConfigApi.md#getdownloadclientconfigbyid) | **Get** /api/v3/config/downloadclient/{id} | 
*DownloadClientConfigApi* | [**UpdateDownloadClientConfig**](whisparr/docs/DownloadClientConfigApi.md#updatedownloadclientconfig) | **Put** /api/v3/config/downloadclient/{id} | 
*ExtraFileApi* | [**ListExtraFile**](whisparr/docs/ExtraFileApi.md#listextrafile) | **Get** /api/v3/extrafile | 
*FileSystemApi* | [**GetFileSystem**](whisparr/docs/FileSystemApi.md#getfilesystem) | **Get** /api/v3/filesystem | 
*FileSystemApi* | [**GetFileSystemMediafiles**](whisparr/docs/FileSystemApi.md#getfilesystemmediafiles) | **Get** /api/v3/filesystem/mediafiles | 
*FileSystemApi* | [**GetFileSystemType**](whisparr/docs/FileSystemApi.md#getfilesystemtype) | **Get** /api/v3/filesystem/type | 
*HealthApi* | [**GetHealthById**](whisparr/docs/HealthApi.md#gethealthbyid) | **Get** /api/v3/health/{id} | 
*HealthApi* | [**ListHealth**](whisparr/docs/HealthApi.md#listhealth) | **Get** /api/v3/health | 
*HistoryApi* | [**CreateHistoryFailedById**](whisparr/docs/HistoryApi.md#createhistoryfailedbyid) | **Post** /api/v3/history/failed/{id} | 
*HistoryApi* | [**GetHistory**](whisparr/docs/HistoryApi.md#gethistory) | **Get** /api/v3/history | 
*HistoryApi* | [**ListHistoryMovie**](whisparr/docs/HistoryApi.md#listhistorymovie) | **Get** /api/v3/history/movie | 
*HistoryApi* | [**ListHistorySince**](whisparr/docs/HistoryApi.md#listhistorysince) | **Get** /api/v3/history/since | 
*HostConfigApi* | [**GetHostConfig**](whisparr/docs/HostConfigApi.md#gethostconfig) | **Get** /api/v3/config/host | 
*HostConfigApi* | [**GetHostConfigById**](whisparr/docs/HostConfigApi.md#gethostconfigbyid) | **Get** /api/v3/config/host/{id} | 
*HostConfigApi* | [**UpdateHostConfig**](whisparr/docs/HostConfigApi.md#updatehostconfig) | **Put** /api/v3/config/host/{id} | 
*ImportExclusionsApi* | [**CreateExclusions**](whisparr/docs/ImportExclusionsApi.md#createexclusions) | **Post** /api/v3/exclusions | 
*ImportExclusionsApi* | [**CreateExclusionsBulk**](whisparr/docs/ImportExclusionsApi.md#createexclusionsbulk) | **Post** /api/v3/exclusions/bulk | 
*ImportExclusionsApi* | [**DeleteExclusions**](whisparr/docs/ImportExclusionsApi.md#deleteexclusions) | **Delete** /api/v3/exclusions/{id} | 
*ImportExclusionsApi* | [**GetExclusionsById**](whisparr/docs/ImportExclusionsApi.md#getexclusionsbyid) | **Get** /api/v3/exclusions/{id} | 
*ImportExclusionsApi* | [**ListExclusions**](whisparr/docs/ImportExclusionsApi.md#listexclusions) | **Get** /api/v3/exclusions | 
*ImportExclusionsApi* | [**UpdateExclusions**](whisparr/docs/ImportExclusionsApi.md#updateexclusions) | **Put** /api/v3/exclusions/{id} | 
*ImportListApi* | [**CreateImportList**](whisparr/docs/ImportListApi.md#createimportlist) | **Post** /api/v3/importlist | 
*ImportListApi* | [**CreateImportListActionByName**](whisparr/docs/ImportListApi.md#createimportlistactionbyname) | **Post** /api/v3/importlist/action/{name} | 
*ImportListApi* | [**DeleteImportList**](whisparr/docs/ImportListApi.md#deleteimportlist) | **Delete** /api/v3/importlist/{id} | 
*ImportListApi* | [**GetImportListById**](whisparr/docs/ImportListApi.md#getimportlistbyid) | **Get** /api/v3/importlist/{id} | 
*ImportListApi* | [**ListImportList**](whisparr/docs/ImportListApi.md#listimportlist) | **Get** /api/v3/importlist | 
*ImportListApi* | [**ListImportListSchema**](whisparr/docs/ImportListApi.md#listimportlistschema) | **Get** /api/v3/importlist/schema | 
*ImportListApi* | [**TestImportList**](whisparr/docs/ImportListApi.md#testimportlist) | **Post** /api/v3/importlist/test | 
*ImportListApi* | [**TestallImportList**](whisparr/docs/ImportListApi.md#testallimportlist) | **Post** /api/v3/importlist/testall | 
*ImportListApi* | [**UpdateImportList**](whisparr/docs/ImportListApi.md#updateimportlist) | **Put** /api/v3/importlist/{id} | 
*ImportListConfigApi* | [**GetImportListConfig**](whisparr/docs/ImportListConfigApi.md#getimportlistconfig) | **Get** /api/v3/config/importlist | 
*ImportListConfigApi* | [**GetImportListConfigById**](whisparr/docs/ImportListConfigApi.md#getimportlistconfigbyid) | **Get** /api/v3/config/importlist/{id} | 
*ImportListConfigApi* | [**UpdateImportListConfig**](whisparr/docs/ImportListConfigApi.md#updateimportlistconfig) | **Put** /api/v3/config/importlist/{id} | 
*ImportListMoviesApi* | [**CreateImportlistMovie**](whisparr/docs/ImportListMoviesApi.md#createimportlistmovie) | **Post** /api/v3/importlist/movie | 
*ImportListMoviesApi* | [**GetImportlistMovie**](whisparr/docs/ImportListMoviesApi.md#getimportlistmovie) | **Get** /api/v3/importlist/movie | 
*IndexerApi* | [**CreateIndexer**](whisparr/docs/IndexerApi.md#createindexer) | **Post** /api/v3/indexer | 
*IndexerApi* | [**CreateIndexerActionByName**](whisparr/docs/IndexerApi.md#createindexeractionbyname) | **Post** /api/v3/indexer/action/{name} | 
*IndexerApi* | [**DeleteIndexer**](whisparr/docs/IndexerApi.md#deleteindexer) | **Delete** /api/v3/indexer/{id} | 
*IndexerApi* | [**GetIndexerById**](whisparr/docs/IndexerApi.md#getindexerbyid) | **Get** /api/v3/indexer/{id} | 
*IndexerApi* | [**ListIndexer**](whisparr/docs/IndexerApi.md#listindexer) | **Get** /api/v3/indexer | 
*IndexerApi* | [**ListIndexerSchema**](whisparr/docs/IndexerApi.md#listindexerschema) | **Get** /api/v3/indexer/schema | 
*IndexerApi* | [**TestIndexer**](whisparr/docs/IndexerApi.md#testindexer) | **Post** /api/v3/indexer/test | 
*IndexerApi* | [**TestallIndexer**](whisparr/docs/IndexerApi.md#testallindexer) | **Post** /api/v3/indexer/testall | 
*IndexerApi* | [**UpdateIndexer**](whisparr/docs/IndexerApi.md#updateindexer) | **Put** /api/v3/indexer/{id} | 
*IndexerConfigApi* | [**GetIndexerConfig**](whisparr/docs/IndexerConfigApi.md#getindexerconfig) | **Get** /api/v3/config/indexer | 
*IndexerConfigApi* | [**GetIndexerConfigById**](whisparr/docs/IndexerConfigApi.md#getindexerconfigbyid) | **Get** /api/v3/config/indexer/{id} | 
*IndexerConfigApi* | [**UpdateIndexerConfig**](whisparr/docs/IndexerConfigApi.md#updateindexerconfig) | **Put** /api/v3/config/indexer/{id} | 
*IndexerFlagApi* | [**ListIndexerFlag**](whisparr/docs/IndexerFlagApi.md#listindexerflag) | **Get** /api/v3/indexerflag | 
*InitializeJsApi* | [**GetInitializeJs**](whisparr/docs/InitializeJsApi.md#getinitializejs) | **Get** /initialize.js | 
*LanguageApi* | [**GetLanguageById**](whisparr/docs/LanguageApi.md#getlanguagebyid) | **Get** /api/v3/language/{id} | 
*LanguageApi* | [**ListLanguage**](whisparr/docs/LanguageApi.md#listlanguage) | **Get** /api/v3/language | 
*LocalizationApi* | [**GetLocalization**](whisparr/docs/LocalizationApi.md#getlocalization) | **Get** /api/v3/localization | 
*LogApi* | [**GetLog**](whisparr/docs/LogApi.md#getlog) | **Get** /api/v3/log | 
*LogFileApi* | [**GetLogFileByFilename**](whisparr/docs/LogFileApi.md#getlogfilebyfilename) | **Get** /api/v3/log/file/{filename} | 
*LogFileApi* | [**ListLogFile**](whisparr/docs/LogFileApi.md#listlogfile) | **Get** /api/v3/log/file | 
*ManualImportApi* | [**CreateManualImport**](whisparr/docs/ManualImportApi.md#createmanualimport) | **Post** /api/v3/manualimport | 
*ManualImportApi* | [**ListManualImport**](whisparr/docs/ManualImportApi.md#listmanualimport) | **Get** /api/v3/manualimport | 
*MediaCoverApi* | [**GetMediaCovermovieIdByFilename**](whisparr/docs/MediaCoverApi.md#getmediacovermovieidbyfilename) | **Get** /api/v3/mediacover/{movieId}/{filename} | 
*MediaManagementConfigApi* | [**GetMediaManagementConfig**](whisparr/docs/MediaManagementConfigApi.md#getmediamanagementconfig) | **Get** /api/v3/config/mediamanagement | 
*MediaManagementConfigApi* | [**GetMediaManagementConfigById**](whisparr/docs/MediaManagementConfigApi.md#getmediamanagementconfigbyid) | **Get** /api/v3/config/mediamanagement/{id} | 
*MediaManagementConfigApi* | [**UpdateMediaManagementConfig**](whisparr/docs/MediaManagementConfigApi.md#updatemediamanagementconfig) | **Put** /api/v3/config/mediamanagement/{id} | 
*MetadataApi* | [**CreateMetadata**](whisparr/docs/MetadataApi.md#createmetadata) | **Post** /api/v3/metadata | 
*MetadataApi* | [**CreateMetadataActionByName**](whisparr/docs/MetadataApi.md#createmetadataactionbyname) | **Post** /api/v3/metadata/action/{name} | 
*MetadataApi* | [**DeleteMetadata**](whisparr/docs/MetadataApi.md#deletemetadata) | **Delete** /api/v3/metadata/{id} | 
*MetadataApi* | [**GetMetadataById**](whisparr/docs/MetadataApi.md#getmetadatabyid) | **Get** /api/v3/metadata/{id} | 
*MetadataApi* | [**ListMetadata**](whisparr/docs/MetadataApi.md#listmetadata) | **Get** /api/v3/metadata | 
*MetadataApi* | [**ListMetadataSchema**](whisparr/docs/MetadataApi.md#listmetadataschema) | **Get** /api/v3/metadata/schema | 
*MetadataApi* | [**TestMetadata**](whisparr/docs/MetadataApi.md#testmetadata) | **Post** /api/v3/metadata/test | 
*MetadataApi* | [**TestallMetadata**](whisparr/docs/MetadataApi.md#testallmetadata) | **Post** /api/v3/metadata/testall | 
*MetadataApi* | [**UpdateMetadata**](whisparr/docs/MetadataApi.md#updatemetadata) | **Put** /api/v3/metadata/{id} | 
*MetadataConfigApi* | [**GetMetadataConfig**](whisparr/docs/MetadataConfigApi.md#getmetadataconfig) | **Get** /api/v3/config/metadata | 
*MetadataConfigApi* | [**GetMetadataConfigById**](whisparr/docs/MetadataConfigApi.md#getmetadataconfigbyid) | **Get** /api/v3/config/metadata/{id} | 
*MetadataConfigApi* | [**UpdateMetadataConfig**](whisparr/docs/MetadataConfigApi.md#updatemetadataconfig) | **Put** /api/v3/config/metadata/{id} | 
*MovieApi* | [**CreateMovie**](whisparr/docs/MovieApi.md#createmovie) | **Post** /api/v3/movie | 
*MovieApi* | [**DeleteMovie**](whisparr/docs/MovieApi.md#deletemovie) | **Delete** /api/v3/movie/{id} | 
*MovieApi* | [**GetMovieById**](whisparr/docs/MovieApi.md#getmoviebyid) | **Get** /api/v3/movie/{id} | 
*MovieApi* | [**ListMovie**](whisparr/docs/MovieApi.md#listmovie) | **Get** /api/v3/movie | 
*MovieApi* | [**UpdateMovie**](whisparr/docs/MovieApi.md#updatemovie) | **Put** /api/v3/movie/{id} | 
*MovieEditorApi* | [**DeleteMovieEditor**](whisparr/docs/MovieEditorApi.md#deletemovieeditor) | **Delete** /api/v3/movie/editor | 
*MovieEditorApi* | [**PutMovieEditor**](whisparr/docs/MovieEditorApi.md#putmovieeditor) | **Put** /api/v3/movie/editor | 
*MovieFileApi* | [**DeleteMovieFile**](whisparr/docs/MovieFileApi.md#deletemoviefile) | **Delete** /api/v3/moviefile/{id} | 
*MovieFileApi* | [**DeleteMovieFileBulk**](whisparr/docs/MovieFileApi.md#deletemoviefilebulk) | **Delete** /api/v3/moviefile/bulk | 
*MovieFileApi* | [**GetMovieFileById**](whisparr/docs/MovieFileApi.md#getmoviefilebyid) | **Get** /api/v3/moviefile/{id} | 
*MovieFileApi* | [**ListMovieFile**](whisparr/docs/MovieFileApi.md#listmoviefile) | **Get** /api/v3/moviefile | 
*MovieFileApi* | [**PutMovieFileEditor**](whisparr/docs/MovieFileApi.md#putmoviefileeditor) | **Put** /api/v3/moviefile/editor | 
*MovieFileApi* | [**UpdateMovieFile**](whisparr/docs/MovieFileApi.md#updatemoviefile) | **Put** /api/v3/moviefile/{id} | 
*MovieImportApi* | [**CreateMovieImport**](whisparr/docs/MovieImportApi.md#createmovieimport) | **Post** /api/v3/movie/import | 
*MovieImportApi* | [**GetMovieImportById**](whisparr/docs/MovieImportApi.md#getmovieimportbyid) | **Get** /api/v3/movie/import/{id} | 
*MovieLookupApi* | [**GetMovieLookup**](whisparr/docs/MovieLookupApi.md#getmovielookup) | **Get** /api/v3/movie/lookup | 
*MovieLookupApi* | [**GetMovieLookupById**](whisparr/docs/MovieLookupApi.md#getmovielookupbyid) | **Get** /api/v3/movie/lookup/{id} | 
*MovieLookupApi* | [**GetMovieLookupImdb**](whisparr/docs/MovieLookupApi.md#getmovielookupimdb) | **Get** /api/v3/movie/lookup/imdb | 
*MovieLookupApi* | [**GetMovieLookupTmdb**](whisparr/docs/MovieLookupApi.md#getmovielookuptmdb) | **Get** /api/v3/movie/lookup/tmdb | 
*NamingConfigApi* | [**GetNamingConfig**](whisparr/docs/NamingConfigApi.md#getnamingconfig) | **Get** /api/v3/config/naming | 
*NamingConfigApi* | [**GetNamingConfigById**](whisparr/docs/NamingConfigApi.md#getnamingconfigbyid) | **Get** /api/v3/config/naming/{id} | 
*NamingConfigApi* | [**GetNamingConfigExamples**](whisparr/docs/NamingConfigApi.md#getnamingconfigexamples) | **Get** /api/v3/config/naming/examples | 
*NamingConfigApi* | [**UpdateNamingConfig**](whisparr/docs/NamingConfigApi.md#updatenamingconfig) | **Put** /api/v3/config/naming/{id} | 
*NotificationApi* | [**CreateNotification**](whisparr/docs/NotificationApi.md#createnotification) | **Post** /api/v3/notification | 
*NotificationApi* | [**CreateNotificationActionByName**](whisparr/docs/NotificationApi.md#createnotificationactionbyname) | **Post** /api/v3/notification/action/{name} | 
*NotificationApi* | [**DeleteNotification**](whisparr/docs/NotificationApi.md#deletenotification) | **Delete** /api/v3/notification/{id} | 
*NotificationApi* | [**GetNotificationById**](whisparr/docs/NotificationApi.md#getnotificationbyid) | **Get** /api/v3/notification/{id} | 
*NotificationApi* | [**ListNotification**](whisparr/docs/NotificationApi.md#listnotification) | **Get** /api/v3/notification | 
*NotificationApi* | [**ListNotificationSchema**](whisparr/docs/NotificationApi.md#listnotificationschema) | **Get** /api/v3/notification/schema | 
*NotificationApi* | [**TestNotification**](whisparr/docs/NotificationApi.md#testnotification) | **Post** /api/v3/notification/test | 
*NotificationApi* | [**TestallNotification**](whisparr/docs/NotificationApi.md#testallnotification) | **Post** /api/v3/notification/testall | 
*NotificationApi* | [**UpdateNotification**](whisparr/docs/NotificationApi.md#updatenotification) | **Put** /api/v3/notification/{id} | 
*ParseApi* | [**GetParse**](whisparr/docs/ParseApi.md#getparse) | **Get** /api/v3/parse | 
*PingApi* | [**GetPing**](whisparr/docs/PingApi.md#getping) | **Get** /ping | 
*QualityDefinitionApi* | [**GetQualityDefinitionById**](whisparr/docs/QualityDefinitionApi.md#getqualitydefinitionbyid) | **Get** /api/v3/qualitydefinition/{id} | 
*QualityDefinitionApi* | [**ListQualityDefinition**](whisparr/docs/QualityDefinitionApi.md#listqualitydefinition) | **Get** /api/v3/qualitydefinition | 
*QualityDefinitionApi* | [**PutQualityDefinitionUpdate**](whisparr/docs/QualityDefinitionApi.md#putqualitydefinitionupdate) | **Put** /api/v3/qualitydefinition/update | 
*QualityDefinitionApi* | [**UpdateQualityDefinition**](whisparr/docs/QualityDefinitionApi.md#updatequalitydefinition) | **Put** /api/v3/qualitydefinition/{id} | 
*QualityProfileApi* | [**CreateQualityProfile**](whisparr/docs/QualityProfileApi.md#createqualityprofile) | **Post** /api/v3/qualityprofile | 
*QualityProfileApi* | [**DeleteQualityProfile**](whisparr/docs/QualityProfileApi.md#deletequalityprofile) | **Delete** /api/v3/qualityprofile/{id} | 
*QualityProfileApi* | [**GetQualityProfileById**](whisparr/docs/QualityProfileApi.md#getqualityprofilebyid) | **Get** /api/v3/qualityprofile/{id} | 
*QualityProfileApi* | [**ListQualityProfile**](whisparr/docs/QualityProfileApi.md#listqualityprofile) | **Get** /api/v3/qualityprofile | 
*QualityProfileApi* | [**UpdateQualityProfile**](whisparr/docs/QualityProfileApi.md#updatequalityprofile) | **Put** /api/v3/qualityprofile/{id} | 
*QualityProfileSchemaApi* | [**GetQualityprofileSchema**](whisparr/docs/QualityProfileSchemaApi.md#getqualityprofileschema) | **Get** /api/v3/qualityprofile/schema | 
*QueueApi* | [**DeleteQueue**](whisparr/docs/QueueApi.md#deletequeue) | **Delete** /api/v3/queue/{id} | 
*QueueApi* | [**DeleteQueueBulk**](whisparr/docs/QueueApi.md#deletequeuebulk) | **Delete** /api/v3/queue/bulk | 
*QueueApi* | [**GetQueue**](whisparr/docs/QueueApi.md#getqueue) | **Get** /api/v3/queue | 
*QueueApi* | [**GetQueueById**](whisparr/docs/QueueApi.md#getqueuebyid) | **Get** /api/v3/queue/{id} | 
*QueueActionApi* | [**CreateQueueGrabBulk**](whisparr/docs/QueueActionApi.md#createqueuegrabbulk) | **Post** /api/v3/queue/grab/bulk | 
*QueueActionApi* | [**CreateQueueGrabById**](whisparr/docs/QueueActionApi.md#createqueuegrabbyid) | **Post** /api/v3/queue/grab/{id} | 
*QueueDetailsApi* | [**GetQueueDetailsById**](whisparr/docs/QueueDetailsApi.md#getqueuedetailsbyid) | **Get** /api/v3/queue/details/{id} | 
*QueueDetailsApi* | [**ListQueueDetails**](whisparr/docs/QueueDetailsApi.md#listqueuedetails) | **Get** /api/v3/queue/details | 
*QueueStatusApi* | [**GetQueueStatus**](whisparr/docs/QueueStatusApi.md#getqueuestatus) | **Get** /api/v3/queue/status | 
*QueueStatusApi* | [**GetQueueStatusById**](whisparr/docs/QueueStatusApi.md#getqueuestatusbyid) | **Get** /api/v3/queue/status/{id} | 
*ReleaseApi* | [**CreateRelease**](whisparr/docs/ReleaseApi.md#createrelease) | **Post** /api/v3/release | 
*ReleaseApi* | [**GetReleaseById**](whisparr/docs/ReleaseApi.md#getreleasebyid) | **Get** /api/v3/release/{id} | 
*ReleaseApi* | [**ListRelease**](whisparr/docs/ReleaseApi.md#listrelease) | **Get** /api/v3/release | 
*ReleasePushApi* | [**CreateReleasePush**](whisparr/docs/ReleasePushApi.md#createreleasepush) | **Post** /api/v3/release/push | 
*ReleasePushApi* | [**GetReleasePushById**](whisparr/docs/ReleasePushApi.md#getreleasepushbyid) | **Get** /api/v3/release/push/{id} | 
*RemotePathMappingApi* | [**CreateRemotePathMapping**](whisparr/docs/RemotePathMappingApi.md#createremotepathmapping) | **Post** /api/v3/remotepathmapping | 
*RemotePathMappingApi* | [**DeleteRemotePathMapping**](whisparr/docs/RemotePathMappingApi.md#deleteremotepathmapping) | **Delete** /api/v3/remotepathmapping/{id} | 
*RemotePathMappingApi* | [**GetRemotePathMappingById**](whisparr/docs/RemotePathMappingApi.md#getremotepathmappingbyid) | **Get** /api/v3/remotepathmapping/{id} | 
*RemotePathMappingApi* | [**ListRemotePathMapping**](whisparr/docs/RemotePathMappingApi.md#listremotepathmapping) | **Get** /api/v3/remotepathmapping | 
*RemotePathMappingApi* | [**UpdateRemotePathMapping**](whisparr/docs/RemotePathMappingApi.md#updateremotepathmapping) | **Put** /api/v3/remotepathmapping/{id} | 
*RenameMovieApi* | [**ListRename**](whisparr/docs/RenameMovieApi.md#listrename) | **Get** /api/v3/rename | 
*RestrictionApi* | [**CreateRestriction**](whisparr/docs/RestrictionApi.md#createrestriction) | **Post** /api/v3/restriction | 
*RestrictionApi* | [**DeleteRestriction**](whisparr/docs/RestrictionApi.md#deleterestriction) | **Delete** /api/v3/restriction/{id} | 
*RestrictionApi* | [**GetRestrictionById**](whisparr/docs/RestrictionApi.md#getrestrictionbyid) | **Get** /api/v3/restriction/{id} | 
*RestrictionApi* | [**ListRestriction**](whisparr/docs/RestrictionApi.md#listrestriction) | **Get** /api/v3/restriction | 
*RestrictionApi* | [**UpdateRestriction**](whisparr/docs/RestrictionApi.md#updaterestriction) | **Put** /api/v3/restriction/{id} | 
*RootFolderApi* | [**CreateRootFolder**](whisparr/docs/RootFolderApi.md#createrootfolder) | **Post** /api/v3/rootfolder | 
*RootFolderApi* | [**DeleteRootFolder**](whisparr/docs/RootFolderApi.md#deleterootfolder) | **Delete** /api/v3/rootfolder/{id} | 
*RootFolderApi* | [**GetRootFolderById**](whisparr/docs/RootFolderApi.md#getrootfolderbyid) | **Get** /api/v3/rootfolder/{id} | 
*RootFolderApi* | [**ListRootFolder**](whisparr/docs/RootFolderApi.md#listrootfolder) | **Get** /api/v3/rootfolder | 
*StaticResourceApi* | [**Get**](whisparr/docs/StaticResourceApi.md#get) | **Get** / | 
*StaticResourceApi* | [**GetByPath**](whisparr/docs/StaticResourceApi.md#getbypath) | **Get** /{path} | 
*StaticResourceApi* | [**GetContentByPath**](whisparr/docs/StaticResourceApi.md#getcontentbypath) | **Get** /content/{path} | 
*StaticResourceApi* | [**GetLogin**](whisparr/docs/StaticResourceApi.md#getlogin) | **Get** /login | 
*SystemApi* | [**CreateSystemRestart**](whisparr/docs/SystemApi.md#createsystemrestart) | **Post** /api/v3/system/restart | 
*SystemApi* | [**CreateSystemShutdown**](whisparr/docs/SystemApi.md#createsystemshutdown) | **Post** /api/v3/system/shutdown | 
*SystemApi* | [**GetSystemRoutes**](whisparr/docs/SystemApi.md#getsystemroutes) | **Get** /api/v3/system/routes | 
*SystemApi* | [**GetSystemRoutesDuplicate**](whisparr/docs/SystemApi.md#getsystemroutesduplicate) | **Get** /api/v3/system/routes/duplicate | 
*SystemApi* | [**GetSystemStatus**](whisparr/docs/SystemApi.md#getsystemstatus) | **Get** /api/v3/system/status | 
*TagApi* | [**CreateTag**](whisparr/docs/TagApi.md#createtag) | **Post** /api/v3/tag | 
*TagApi* | [**DeleteTag**](whisparr/docs/TagApi.md#deletetag) | **Delete** /api/v3/tag/{id} | 
*TagApi* | [**GetTagById**](whisparr/docs/TagApi.md#gettagbyid) | **Get** /api/v3/tag/{id} | 
*TagApi* | [**ListTag**](whisparr/docs/TagApi.md#listtag) | **Get** /api/v3/tag | 
*TagApi* | [**UpdateTag**](whisparr/docs/TagApi.md#updatetag) | **Put** /api/v3/tag/{id} | 
*TagDetailsApi* | [**GetTagDetailById**](whisparr/docs/TagDetailsApi.md#gettagdetailbyid) | **Get** /api/v3/tag/detail/{id} | 
*TagDetailsApi* | [**ListTagDetail**](whisparr/docs/TagDetailsApi.md#listtagdetail) | **Get** /api/v3/tag/detail | 
*TaskApi* | [**GetSystemTaskById**](whisparr/docs/TaskApi.md#getsystemtaskbyid) | **Get** /api/v3/system/task/{id} | 
*TaskApi* | [**ListSystemTask**](whisparr/docs/TaskApi.md#listsystemtask) | **Get** /api/v3/system/task | 
*UiConfigApi* | [**GetUiConfig**](whisparr/docs/UiConfigApi.md#getuiconfig) | **Get** /api/v3/config/ui | 
*UiConfigApi* | [**GetUiConfigById**](whisparr/docs/UiConfigApi.md#getuiconfigbyid) | **Get** /api/v3/config/ui/{id} | 
*UiConfigApi* | [**UpdateUiConfig**](whisparr/docs/UiConfigApi.md#updateuiconfig) | **Put** /api/v3/config/ui/{id} | 
*UpdateApi* | [**ListUpdate**](whisparr/docs/UpdateApi.md#listupdate) | **Get** /api/v3/update | 
*UpdateLogFileApi* | [**GetLogFileUpdateByFilename**](whisparr/docs/UpdateLogFileApi.md#getlogfileupdatebyfilename) | **Get** /api/v3/log/file/update/{filename} | 
*UpdateLogFileApi* | [**ListLogFileUpdate**](whisparr/docs/UpdateLogFileApi.md#listlogfileupdate) | **Get** /api/v3/log/file/update | 


## Documentation For Models

 - [AddMovieMethod](docs/AddMovieMethod.md)
 - [AddMovieOptions](docs/AddMovieOptions.md)
 - [AlternativeTitle](docs/AlternativeTitle.md)
 - [AlternativeTitleResource](docs/AlternativeTitleResource.md)
 - [ApiInfoResource](docs/ApiInfoResource.md)
 - [ApplyTags](docs/ApplyTags.md)
 - [AuthenticationType](docs/AuthenticationType.md)
 - [BackupResource](docs/BackupResource.md)
 - [BackupType](docs/BackupType.md)
 - [BlocklistBulkResource](docs/BlocklistBulkResource.md)
 - [BlocklistResource](docs/BlocklistResource.md)
 - [BlocklistResourcePagingResource](docs/BlocklistResourcePagingResource.md)
 - [CertificateValidationType](docs/CertificateValidationType.md)
 - [CollectionMovieResource](docs/CollectionMovieResource.md)
 - [CollectionResource](docs/CollectionResource.md)
 - [CollectionUpdateResource](docs/CollectionUpdateResource.md)
 - [ColonReplacementFormat](docs/ColonReplacementFormat.md)
 - [Command](docs/Command.md)
 - [CommandPriority](docs/CommandPriority.md)
 - [CommandResource](docs/CommandResource.md)
 - [CommandStatus](docs/CommandStatus.md)
 - [CommandTrigger](docs/CommandTrigger.md)
 - [CreditResource](docs/CreditResource.md)
 - [CreditType](docs/CreditType.md)
 - [CustomFilterResource](docs/CustomFilterResource.md)
 - [CustomFormatResource](docs/CustomFormatResource.md)
 - [CustomFormatSpecificationSchema](docs/CustomFormatSpecificationSchema.md)
 - [DatabaseType](docs/DatabaseType.md)
 - [DelayProfileResource](docs/DelayProfileResource.md)
 - [DiskSpaceResource](docs/DiskSpaceResource.md)
 - [DownloadClientConfigResource](docs/DownloadClientConfigResource.md)
 - [DownloadClientResource](docs/DownloadClientResource.md)
 - [DownloadProtocol](docs/DownloadProtocol.md)
 - [ExtraFileResource](docs/ExtraFileResource.md)
 - [ExtraFileType](docs/ExtraFileType.md)
 - [Field](docs/Field.md)
 - [FileDateType](docs/FileDateType.md)
 - [HealthCheckResult](docs/HealthCheckResult.md)
 - [HealthResource](docs/HealthResource.md)
 - [HistoryResource](docs/HistoryResource.md)
 - [HistoryResourcePagingResource](docs/HistoryResourcePagingResource.md)
 - [HostConfigResource](docs/HostConfigResource.md)
 - [ImportExclusionsResource](docs/ImportExclusionsResource.md)
 - [ImportListConfigResource](docs/ImportListConfigResource.md)
 - [ImportListResource](docs/ImportListResource.md)
 - [ImportListType](docs/ImportListType.md)
 - [IndexerConfigResource](docs/IndexerConfigResource.md)
 - [IndexerFlagResource](docs/IndexerFlagResource.md)
 - [IndexerResource](docs/IndexerResource.md)
 - [Language](docs/Language.md)
 - [LanguageResource](docs/LanguageResource.md)
 - [LogFileResource](docs/LogFileResource.md)
 - [LogResource](docs/LogResource.md)
 - [LogResourcePagingResource](docs/LogResourcePagingResource.md)
 - [ManualImportReprocessResource](docs/ManualImportReprocessResource.md)
 - [ManualImportResource](docs/ManualImportResource.md)
 - [MediaCover](docs/MediaCover.md)
 - [MediaCoverTypes](docs/MediaCoverTypes.md)
 - [MediaInfoResource](docs/MediaInfoResource.md)
 - [MediaManagementConfigResource](docs/MediaManagementConfigResource.md)
 - [MetadataConfigResource](docs/MetadataConfigResource.md)
 - [MetadataResource](docs/MetadataResource.md)
 - [Modifier](docs/Modifier.md)
 - [MonitorTypes](docs/MonitorTypes.md)
 - [MovieCollection](docs/MovieCollection.md)
 - [MovieEditorResource](docs/MovieEditorResource.md)
 - [MovieFileListResource](docs/MovieFileListResource.md)
 - [MovieFileResource](docs/MovieFileResource.md)
 - [MovieHistoryEventType](docs/MovieHistoryEventType.md)
 - [MovieMetadata](docs/MovieMetadata.md)
 - [MovieResource](docs/MovieResource.md)
 - [MovieRuntimeFormatType](docs/MovieRuntimeFormatType.md)
 - [MovieStatusType](docs/MovieStatusType.md)
 - [MovieTranslation](docs/MovieTranslation.md)
 - [NamingConfigResource](docs/NamingConfigResource.md)
 - [NotificationResource](docs/NotificationResource.md)
 - [PagingResourceFilter](docs/PagingResourceFilter.md)
 - [ParseResource](docs/ParseResource.md)
 - [ParsedMovieInfo](docs/ParsedMovieInfo.md)
 - [PingResource](docs/PingResource.md)
 - [ProfileFormatItemResource](docs/ProfileFormatItemResource.md)
 - [ProperDownloadTypes](docs/ProperDownloadTypes.md)
 - [ProviderMessage](docs/ProviderMessage.md)
 - [ProviderMessageType](docs/ProviderMessageType.md)
 - [ProxyType](docs/ProxyType.md)
 - [Quality](docs/Quality.md)
 - [QualityDefinitionResource](docs/QualityDefinitionResource.md)
 - [QualityModel](docs/QualityModel.md)
 - [QualityProfileQualityItemResource](docs/QualityProfileQualityItemResource.md)
 - [QualityProfileResource](docs/QualityProfileResource.md)
 - [QueueBulkResource](docs/QueueBulkResource.md)
 - [QueueResource](docs/QueueResource.md)
 - [QueueResourcePagingResource](docs/QueueResourcePagingResource.md)
 - [QueueStatusResource](docs/QueueStatusResource.md)
 - [RatingChild](docs/RatingChild.md)
 - [RatingType](docs/RatingType.md)
 - [Ratings](docs/Ratings.md)
 - [Rejection](docs/Rejection.md)
 - [RejectionType](docs/RejectionType.md)
 - [ReleaseResource](docs/ReleaseResource.md)
 - [RemotePathMappingResource](docs/RemotePathMappingResource.md)
 - [RenameMovieResource](docs/RenameMovieResource.md)
 - [RescanAfterRefreshType](docs/RescanAfterRefreshType.md)
 - [RestrictionResource](docs/RestrictionResource.md)
 - [Revision](docs/Revision.md)
 - [RootFolderResource](docs/RootFolderResource.md)
 - [RuntimeMode](docs/RuntimeMode.md)
 - [SelectOption](docs/SelectOption.md)
 - [SortDirection](docs/SortDirection.md)
 - [Source](docs/Source.md)
 - [SourceType](docs/SourceType.md)
 - [SystemResource](docs/SystemResource.md)
 - [TMDbCountryCode](docs/TMDbCountryCode.md)
 - [TagDetailsResource](docs/TagDetailsResource.md)
 - [TagResource](docs/TagResource.md)
 - [TaskResource](docs/TaskResource.md)
 - [TrackedDownloadState](docs/TrackedDownloadState.md)
 - [TrackedDownloadStatus](docs/TrackedDownloadStatus.md)
 - [TrackedDownloadStatusMessage](docs/TrackedDownloadStatusMessage.md)
 - [UiConfigResource](docs/UiConfigResource.md)
 - [UnmappedFolder](docs/UnmappedFolder.md)
 - [UpdateChanges](docs/UpdateChanges.md)
 - [UpdateMechanism](docs/UpdateMechanism.md)
 - [UpdateResource](docs/UpdateResource.md)


## Documentation For Authorization



### X-Api-Key

- **Type**: API key
- **API key parameter name**: X-Api-Key
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: X-Api-Key and passed in as the auth context for each request.


### apikey

- **Type**: API key
- **API key parameter name**: apikey
- **Location**: URL query string

Note, each API key must be added to a map of `map[string]APIKey` where the key is: apikey and passed in as the auth context for each request.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



