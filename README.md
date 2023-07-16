# gin-template
目前這包專案使用到的package如下
1. github.com/spf13/cobra
2. go.uber.org/zap
3. github.com/gin-contrib/zap 
4. github.com/gin-contrib/pprof
<hr />


## github.com/spf13/cobra
Cobra 是一個用於建立命令列應用程式 (CLI) 的 Go 語言套件。這兩個套件通常一起使用的原因如下：<br />
提供 CLI 控制：當您使用 Gin 構建 Web 應用程式時，可能需要運行一些 CLI 命令來管理應用程式，例如啟動/停止服務、檢查配置、創建資料庫等。<br />
透過 Cobra，您可以輕鬆地定義和管理這些 CLI 命令，使得您的 Web 應用程式更具彈性和易用性。<br />

結合配置選項：Cobra 提供了一個簡單而強大的方式來處理命令列選項和參數。這使得您可以使用 CLI 命令來傳遞配置選項給 Gin 應用程式，而不需要硬編碼這些選項。這使得配置管理更加容易且靈活。<br />

建立工具和指令：除了管理 Web 應用程式外，您可能還需要建立一些自定義工具或指令行程式。Cobra 提供了組織和管理這些工具和指令的架構，可以幫助您構建更具組織性和可擴展性的應用程式。 <br />
https://dusty-aunt-c76.notion.site/cobra-5538672d334d4ea3b1932b46fa751942

## go.uber.org/zap
go.uber.org/zap 是 Uber 公司開源的 Go 語言日誌記錄（logging）庫。它的目標是幫助 Go 開發人員在應用程式中實現高性能且結構化的日誌記錄。<br />

zap 被設計成快速且低延遲，同時保持日誌記錄內容的結構化，以方便後續的日誌處理和分析。它提供多種日誌級別，允許開發人員根據需要選擇不同的日誌級別，
例如 Debug、Info、Warning 和 Error 等。<br />

這個庫提供了簡單易用的API，讓開發人員可以輕鬆地記錄日誌，並可以通過不同的輸出方式（例如控制台、文件、網路等）來配置日誌的目的地。

## github.com/gin-contrib/zap 
這是Gin框架的一個中間件（middleware），它允許您在使用Gin框架時，將請求和回應的紀錄進行結構化的日誌記錄。這個中間件使用了go.uber.org/zap庫來處理日誌記錄。<br />

在使用github.com/gin-contrib/zap時，您需要先引入 go.uber.org/zap 庫作為它的依賴。通常，這是在您的Go模塊的go.mod文件中進行設置，這樣當您編譯或運行代碼時，Go語言的模塊系統會自動下載所需的依賴。

## github.com/gin-contrib/pprof
pprof 是 Go 語言的性能分析工具，用於幫助開發人員評估和優化應用程序的性能。<br />

提供了一個方便的方式將 pprof 整合到 Gin Web 框架中。通過引入這個套件，您可以在 Gin 應用程序中訪問 pprof 相關的端點，並在瀏覽器中使用網頁界面來查看性能分析數據。<br />
這有助於您更好地理解應用程序的運行時行為，並找出可能的性能瓶頸，以便進行優化。<br />

查看性能分析內容：<br />
現在您的應用程序應該正在運行並聽取端口 8080。您可以在瀏覽器中訪問以下 URL 來查看性能分析內容：<br />

CPU 分析: {{domain}}/debug/pprof/profile<br />
內存分析: {{domain}}/debug/pprof/heap<br />
goroutine 數量: {{domain}}/debug/pprof/goroutine<br />
互斥鎖競爭情況: {{domain}}/debug/pprof/mutex<br />
瀏覽這些 URL，您應該能夠看到性能分析數據，並使用 pprof 的網頁界面來分析應用程序的性能。這些數據將幫助您了解應用程序的運行狀況，找出潛在的性能問題，以便進行優化。
