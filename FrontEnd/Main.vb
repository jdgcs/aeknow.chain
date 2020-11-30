Imports System.IO
Imports Newtonsoft.Json
Imports Newtonsoft.Json.Linq
Public Class Main
    Public Structure Pointer
        Dim id As String
        Dim key As String
    End Structure

    Private Sub AENSAddress_KeyPress(sender As Object, e As KeyPressEventArgs) Handles AENSAddress.KeyPress
        If e.KeyChar = Convert.ToChar(13) Then
            ' MsgBox(aensAddress.Text)
            If InStr(AENSAddress.Text, "ttp://") > 0 Or InStr(AENSAddress.Text, "ttps://") > 0 Or InStr(AENSAddress.Text, "ww.") > 0 Then
                WebBrowser1.Navigate(AENSAddress.Text)
            End If

            If InStr(AENSAddress.Text, ".test") > 0 Or InStr(AENSAddress.Text, ".chain") > 0 Then
                viewAENS(AENSAddress.Text)
            End If
        End If
    End Sub

    Private Sub ViewAENS(aensAddress)
        '解析AENS，优先访问IPFS还是IPNS？或者解析IPNS后对应。
        Dim AENSStr As String
        Dim IPFSAddress, IPNSAddress As String
        Dim URLAddress As String = ""
        Dim PublicNode As String = Trim(GetJsonConfig(Application.StartupPath & "/data/config.json", "PublicNode"))
        Dim IPFSNode As String = Trim(GetJsonConfig(Application.StartupPath & "/data/config.json", "IPFSNode"))
        AENSStr = WRequest(PublicNode & "/v2/names/" & aensAddress, "GET", "")
        If InStr(AENSStr, "pointers") > 0 Then
            IPFSAddress = GetIPFSAddress(AENSStr)
            IPNSAddress = GetIPNSAddress(AENSStr)

            If IPFSAddress <> "" Then
                URLAddress = IPFSNode & "/ipfs/" & IPFSAddress
            Else
                If IPNSAddress <> "" Then
                    URLAddress = IPFSNode & "/ipns/" & IPNSAddress
                Else
                    MsgBox("Can not resolve IPFS/IPNS")
                End If
            End If
        Else
            MsgBox("Can not resolve " & aensAddress & "Return:" & AENSStr)
        End If

        WebBrowser1.Navigate(URLAddress)
    End Sub

    Private Sub Button1_Click(sender As Object, e As EventArgs) Handles Button1.Click
        AEKPlayer.Show()
        AEKPlayer.PlayVideo(AENSAddress.Text, "url")
    End Sub

    Private Sub NotifyIcon1_MouseClick(sender As Object, e As MouseEventArgs) Handles NotifyIcon1.MouseClick
        If e.Button = Windows.Forms.MouseButtons.Left Then
            TrayMenu.Enabled = False
            If Me.WindowState = FormWindowState.Minimized Then
                Me.WindowState = FormWindowState.Normal '还原
            End If
            Me.Show()
            Me.Activate()
            'Me.NotifyIcon1.Visible = False
            Me.ShowInTaskbar = True
        Else
            TrayMenu.Enabled = True
        End If
    End Sub

    Private Sub Main_Load(sender As Object, e As EventArgs) Handles MyBase.Load
        StartServer()
        WebBrowser1.Navigate("http://127.0.0.1:8888")
        InitBookmarks()
    End Sub
    Private Sub InitBookmarks()
        '初始化菜单
        Dim doItem As ToolStripMenuItem = ToolStripButton4.DropDownItems("Bookmark")
        doItem.DropDownItems.Add("0100", Nothing, AddressOf Bookmark_Click).Name = "idx0"
        doItem.DropDownItems.Add("111", Nothing, AddressOf Bookmark_Click).Name = "idx1"
    End Sub
    Private Sub Bookmark_Click(sender As Object, e As EventArgs) Handles Bookmark.Click


        Try
            Select Case sender.name
                Case "idx0"
                    MsgBox("文件")
                Case "idx1"
                    MsgBox("编辑")
                Case "mnuOpen"
                    MsgBox("打开")
                Case "mnuSave"
                    MsgBox("保存")

            End Select

        Catch ex As Exception

        End Try

    End Sub
    Private Sub StopServer()
        Shell("taskkill /IM v1.exe /F", AppWinStyle.Hide)
        Shell("taskkill /IM ipfs.exe /F", AppWinStyle.Hide)
    End Sub

    Private Sub StartServer()
        Dim proc As New Process

        proc.StartInfo.UseShellExecute = False
        proc.StartInfo.RedirectStandardInput = True
        proc.StartInfo.RedirectStandardOutput = True
        proc.StartInfo.RedirectStandardError = True
        proc.StartInfo.CreateNoWindow = True
        proc.StartInfo.WindowStyle = ProcessWindowStyle.Hidden

        proc.StartInfo.FileName = "v1.exe"

        ' proc.Start()


    End Sub
    Private Sub ToolStripButton1_Click(sender As Object, e As EventArgs) Handles ToolStripButton1.Click
        WebBrowser1.GoBack()
    End Sub

    Private Sub ToolStripButton2_Click(sender As Object, e As EventArgs) Handles ToolStripButton2.Click
        WebBrowser1.GoForward()
    End Sub

    Private Sub ToolStripButton3_Click(sender As Object, e As EventArgs) Handles ToolStripButton3.Click
        WebBrowser1.Refresh()
    End Sub

    Private Sub UpdateNodeInfo()
        Dim UrlAddress As String
        Dim sc As New MSScriptControl.ScriptControl
        Dim JsonStr, IPFSInfo As String
        IPFSInfo = "IPFS:"
        UrlAddress = "http://127.0.0.1:5001/api/v0/version"
        JsonStr = WRequest(UrlAddress, "POST", "")

        If InStr(JsonStr, "Version") Then
            sc.Language = "Jscript"
            sc.AddCode("var js=" & JsonStr)
            IPFSInfo = IPFSInfo & sc.Eval("js.Version")
            ' MsgBox(sc.Eval("js.dataa.ppiUsername"))
            'Timer1.Enabled = False
        Else
            IPFSInfo = IPFSInfo & "OFFLINE"
        End If

        ToolStripStatusLabel1.Text = IPFSInfo
    End Sub
    Private Function GetJsonConfig(JsonFile As String, TargetItem As String) As String
        Dim sc As New MSScriptControl.ScriptControl
        Dim JsonStr As String
        JsonStr = ""

        Dim sr As StreamReader = New StreamReader(JsonFile, System.Text.Encoding.Default)
        Do While sr.Peek() > 0
            JsonStr = JsonStr + sr.ReadLine()
        Loop
        sr.Close()
        sr = Nothing

        If InStr(JsonStr, "}") Then
            sc.Language = "Jscript"
            sc.AddCode("var js=" & JsonStr)
            Return sc.Eval("js." & TargetItem)
        Else
            Return "ERROR:NO JSON STRING"
        End If
        Return "NULL"
    End Function
    Private Function GetJsonItem(JsonStr As String, TargetItem As String) As String
        Dim sc As New MSScriptControl.ScriptControl
        If InStr(JsonStr, "}") Then
            sc.Language = "Jscript"
            sc.AddCode("var js=" & JsonStr)
            Return sc.Eval("js." & TargetItem)
        Else
            Return "ERROR:NO JSON STRING"
        End If
        Return "NULL"
    End Function

    Public Function GetIPNSAddress(AENSStr As String) As String
        If InStr(AENSStr, "ak_ipnsoMiJmYq1joKGXFtLRDrSJ3mUjapNB7gcPud7mmpVUXssM") Then
            Dim p As JObject = CType(JsonConvert.DeserializeObject(AENSStr), JObject)
            Dim Mypointer As List(Of Pointer)
            Mypointer = JsonConvert.DeserializeObject(Of List(Of Pointer))(p("pointers").ToString)

            Dim pointerCount As Integer = Mypointer.Count
            If pointerCount > 0 Then
                For i = 0 To pointerCount
                    If Mypointer(i).id = "ak_ipnsoMiJmYq1joKGXFtLRDrSJ3mUjapNB7gcPud7mmpVUXssM" Then
                        Return Mypointer(i).key
                    End If
                Next
            End If
        End If
        Return ""
    End Function
    Public Function GetIPFSAddress(AENSStr As String) As String
        If InStr(AENSStr, "ak_ipfsD1iUfRLdnJjQMEczjSzzphPbNnSQudnqUAe1vPJetmMK9") Then
            Dim p As JObject = CType(JsonConvert.DeserializeObject(AENSStr), JObject)
            Dim Mypointer As List(Of Pointer)
            Mypointer = JsonConvert.DeserializeObject(Of List(Of Pointer))(p("pointers").ToString)

            Dim pointerCount As Integer = Mypointer.Count
            If pointerCount > 0 Then
                For i = 0 To pointerCount
                    If Mypointer(i).id = "ak_ipfsD1iUfRLdnJjQMEczjSzzphPbNnSQudnqUAe1vPJetmMK9" Then
                        Return Mypointer(i).key
                    End If
                Next
            End If
        End If
        Return ""
    End Function
    Function WRequest(URL As String, method As String, POSTdata As String) As String
        Dim responseData As String = ""
        Try
            Dim cookieJar As New Net.CookieContainer()
            Dim hwrequest As Net.HttpWebRequest = Net.WebRequest.Create(URL)
            hwrequest.CookieContainer = cookieJar
            hwrequest.Accept = "*/*"
            hwrequest.AllowAutoRedirect = True
            hwrequest.UserAgent = "AEKnow.chain/Alpha"
            hwrequest.Timeout = 60000
            hwrequest.Method = method
            If hwrequest.Method = "POST" Then
                hwrequest.ContentType = "application/x-www-form-urlencoded"
                Dim encoding As New Text.ASCIIEncoding() 'Use UTF8Encoding for XML requests
                Dim postByteArray() As Byte = encoding.GetBytes(POSTdata)
                hwrequest.ContentLength = postByteArray.Length
                Dim postStream As IO.Stream = hwrequest.GetRequestStream()
                postStream.Write(postByteArray, 0, postByteArray.Length)
                postStream.Close()
            End If
            Dim hwresponse As Net.HttpWebResponse = hwrequest.GetResponse()
            If hwresponse.StatusCode = Net.HttpStatusCode.OK Then
                Dim responseStream As IO.StreamReader =
                  New IO.StreamReader(hwresponse.GetResponseStream())
                responseData = responseStream.ReadToEnd()
            End If
            hwresponse.Close()
        Catch e As Exception
            responseData = "An error occurred: " & e.Message
        End Try
        Return responseData
    End Function

    Private Sub AENSAddress_Click(sender As Object, e As EventArgs) Handles AENSAddress.Click

    End Sub

    Private Sub Timer1_Tick(sender As Object, e As EventArgs) Handles Timer1.Tick
        UpdateNodeInfo()
    End Sub

    Private Sub ToolStripButton5_Click(sender As Object, e As EventArgs) Handles ToolStripButton5.Click
        Dim HomeAddress As String
        HomeAddress = "http://127.0.0.1:8888"
        WebBrowser1.Navigate(HomeAddress)
    End Sub

    Private Sub WebBrowser1_Navigating(sender As Object, e As WebBrowserNavigatingEventArgs) Handles WebBrowser1.Navigating
        '  MsgBox(e.Url.ToString)
        If InStr(e.Url.ToString, "/ipfs/Qm") > 1 And InStr(e.Url.ToString, "?video") > 1 Then
            AEKPlayer.Show()
            AEKPlayer.PlayVideo(AENSAddress.Text, "url")
            e.Cancel = True
        End If
    End Sub

    Private Sub Main_Resize(sender As Object, e As EventArgs) Handles Me.Resize
        If (Me.Width - 100) > 0 Then
            WebBrowser1.Width = Me.Width - 24
        Else
            ' WebBrowser1.Width = 0
        End If
        'TopNavBar.ae
        AENSAddress.Width = Me.Width - 500

        If (Me.Height - 100) > 0 Then
            WebBrowser1.Height = Me.Height - 88
        Else
            'WebBrowser1.Height = 0
        End If
    End Sub
End Class
