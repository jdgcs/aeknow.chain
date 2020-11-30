Imports System.IO
Imports LibVLCSharp.Shared
'Imports VLC
Imports LibVLCSharp.WinForms
Imports System.Runtime.InteropServices
Imports System.ComponentModel

Public Class AEKPlayer
    Private fullnow As Boolean
    Private rect As Rectangle
    Public Const SPI_GETWORKAREA As Integer = &H30
    Public Const SPI_SETWORKAREA As Integer = &H2F
    Public Const SPIF_UPDATEINIFILE As Integer = 1
    Public Const SW_HIDE As Integer = 0
    Public Const SW_SHOW As Integer = 5

    <DllImport("user32.dll")>
    Private Shared Function FindWindow(ByVal lpClassName As String, ByVal lpWindowName As String) As Integer

    End Function
    <DllImport("user32.dll")>
    Public Shared Function ShowWindow(ByVal hwnd As Integer, ByVal nCmdShow As Integer) As Integer

    End Function
    <DllImport("user32.dll")>
    Private Shared Function SystemParametersInfo(ByVal uAction As Integer, ByVal uParam As Integer, ByRef lpvParam As Rectangle, ByVal fuWinIni As Integer) As Integer

    End Function

    ''' <summary>
    ''' fullscreen'''''''''''''''''''''''
    ''' </summary>
    Public video As LibVLC
    Public mp As MediaPlayer
    Public IsActivated As Boolean = True '当前播放器是否聚焦，不聚焦的话控制栏不显示
    Public IsShowToolTip As Boolean = False
    Public VideoLength As Integer = 0
    Public HoverSecs As Integer = 0
    Protected Overrides Sub OnDoubleClick(ByVal e As EventArgs)
        MyBase.OnDoubleClick(e)
        Me.fullnow = Not Me.fullnow
        AEKPlayer.SetFullScreen(Me.fullnow, (Me.rect))
        MyBase.FormBorderStyle = IIf(Me.fullnow, FormBorderStyle.None, FormBorderStyle.FixedDialog)
        MyBase.WindowState = IIf(Me.fullnow, FormWindowState.Maximized, FormWindowState.Normal)
    End Sub
    Public Shared Function SetFullScreen(ByVal fullscreen As Boolean, ByRef rectOld As Rectangle) As Boolean
        Dim Hwnd As Integer = AEKPlayer.FindWindow("Shell_TrayWnd", Nothing)
        If (Hwnd = 0) Then
            Return False
        End If
        If fullscreen Then
            AEKPlayer.ShowWindow(Hwnd, 0)
            Dim rectFull As Rectangle = Screen.PrimaryScreen.Bounds
            AEKPlayer.SystemParametersInfo(&H30, 0, (rectOld), 1)
            AEKPlayer.SystemParametersInfo(&H2F, 0, (rectFull), 1)
        Else
            AEKPlayer.ShowWindow(Hwnd, 5)
            AEKPlayer.SystemParametersInfo(&H2F, 0, (rectOld), 1)
        End If
        Return True
    End Function
    Private Sub AEKPlayer_Load(sender As Object, e As EventArgs) Handles MyBase.Load
        Label1.Text = Cursor.Position.X
    End Sub

    Private Sub AEKPlayer_Resize(sender As Object, e As EventArgs) Handles Me.Resize
        'AEKPanel.Bottom = 0
    End Sub

    Private Sub AEKPlay_Click(sender As Object, e As EventArgs) Handles AEKPlay.Click
        Dim VideoAddress As String = "http://127.0.0.1:8080/ipfs/Qmf2y99mayUJJFFrHFjD2h9yLcDPc6sD3FAgVSmpCq5xMK/"
        ' VideoAddress = "http://127.0.0.1:8080/ipfs/QmfJcKLZzdX99xt7cXfhMa2Pf9PBasEeFSqYiRDu73ovrG/"
        ' VideoAddress = "d:\reach.wmv"
        ' PlayVideo(VideoAddress, "file")
        PlayVideo(VideoAddress, "url")
    End Sub

    Public Sub PlayVideo(ByVal VideoAddress As String, vType As String)

        If AEKPlay.ToolTipText = "Play/Open" Then
            Core.Initialize()

            Using libVLC = New LibVLC()

                If vType = "url" Then
                    Dim video = New Media(libVLC, VideoAddress, FromType.FromLocation)
                    Dim mp As New MediaPlayer(video)
                    AEKView.MediaPlayer = mp
                    mp.Play()

                Else
                    Dim video = New Media(libVLC, VideoAddress, FromType.FromPath)
                    Dim mp As New MediaPlayer(video)
                    AEKView.MediaPlayer = mp
                    mp.Play()
                End If

            End Using

            TrackBar1.Maximum = 100 '7987502
            Timer1.Enabled = True

            AEKPlay.Image = ImageList1.Images(5)
            AEKPlay.ToolTipText = "Pause"
        ElseIf AEKPlay.ToolTipText = "Pause" Then
            AEKView.MediaPlayer.Pause()
            AEKPlay.ToolTipText = "Resume"
            AEKPlay.Image = ImageList1.Images(6)
        ElseIf AEKPlay.ToolTipText = "Resume" Then
            AEKView.MediaPlayer.Play()
            AEKPlay.Image = ImageList1.Images(5)
            AEKPlay.ToolTipText = "Pause"
        End If
    End Sub

    Private Sub AEKStop_Click(sender As Object, e As EventArgs) Handles AEKStop.Click
        If AEKView.MediaPlayer Is Nothing Then
            MsgBox("NULL")
        Else
            AEKView.MediaPlayer.Stop()
            AEKPlay.ToolTipText = "Play/Open"
            AEKPlay.Image = ImageList1.Images(6)
        End If
    End Sub

    Private Sub AEKScreen_Click(sender As Object, e As EventArgs) Handles AEKScreen.Click
        MyBase.OnDoubleClick(e)
        Me.fullnow = Not Me.fullnow
        AEKPlayer.SetFullScreen(Me.fullnow, (Me.rect))
        MyBase.FormBorderStyle = IIf(Me.fullnow, FormBorderStyle.None, FormBorderStyle.FixedDialog)
        MyBase.WindowState = IIf(Me.fullnow, FormWindowState.Maximized, FormWindowState.Normal)

        If AEKScreen.ToolTipText = "FullScreen" Then
            AEKScreen.Image = ImageList1.Images(8)
            AEKScreen.ToolTipText = "Exit FullScreen"
        Else
            AEKScreen.Image = ImageList1.Images(8)
            AEKScreen.ToolTipText = "FullScreen"
        End If
    End Sub

    Private Sub Timer1_Tick(sender As Object, e As EventArgs) Handles Timer1.Tick
        Dim pubfunc As New PubFuncs

        If AEKView.MediaPlayer.Length > 0 Then '如果视频的长度存在>0,开始读取进度等操作
            TrackBar1.Value = Int((AEKView.MediaPlayer.Time / AEKView.MediaPlayer.Length) * 100)
            AEKCurrent.Text = pubfunc.ZToFSAll(Int(AEKView.MediaPlayer.Time / 1000))
            AEKFull.Text = pubfunc.ZToFSAll(Int(AEKView.MediaPlayer.Length / 1000))
            VideoLength = Int(AEKView.MediaPlayer.Length / 1000)

            If Label1.Text <> Cursor.Position.X Then
                TimerBar.Enabled = True
                If IsActivated Then
                    AEKPanel.Visible = True
                End If
            End If

            Label1.Text = Cursor.Position.X
        End If
    End Sub

    Private Sub TimerBar_Tick(sender As Object, e As EventArgs) Handles TimerBar.Tick
        AEKPanel.Visible = False
        TimerBar.Enabled = False
        IsShowToolTip = False
        ToolTip1.Hide(Me.TrackBar1)
    End Sub


    Private Sub AEKPlayer_Activated(sender As Object, e As EventArgs) Handles Me.Activated
        IsActivated = True
    End Sub

    Private Sub AEKPlayer_Deactivate(sender As Object, e As EventArgs) Handles Me.Deactivate
        IsActivated = False
    End Sub

    Private Sub TrackBar1_MouseMove(sender As Object, e As MouseEventArgs) Handles TrackBar1.MouseMove
        IsShowToolTip = True
        Dim pubfunc As New PubFuncs
        Dim borderW As Integer = 12
        Dim barLen As Integer = TrackBar1.Width - borderW
        Dim curPos As Integer = e.X - borderW / 2
        If curPos > barLen Then curPos = barLen
        If curPos < 0 Then curPos = 0

        If Cursor.Position.X - Me.TrackBar1.Location.X > 0 Then
            HoverSecs = Int((curPos / barLen) * VideoLength)
            If IsShowToolTip Then
                ToolTip1.Show(pubfunc.ZToFSAll(HoverSecs), Me.TrackBar1, e.X, e.Y - 20)
            End If
        End If
    End Sub


    Private Sub TrackBar1_MouseHover(sender As Object, e As EventArgs) Handles TrackBar1.MouseHover
        'MsgBox("hi")
    End Sub

    Private Sub TrackBar1_MouseUp(sender As Object, e As MouseEventArgs) Handles TrackBar1.MouseUp
        Dim pubfunc As New PubFuncs
        Dim borderW As Integer = 12
        Dim barLen As Integer = TrackBar1.Width - borderW
        Dim curPos As Integer = e.X - borderW / 2
        If curPos > barLen Then curPos = barLen
        If curPos < 0 Then curPos = 0

        TrackBar1.Value = Int(curPos * TrackBar1.Maximum / barLen)
        AEKView.MediaPlayer.Time = HoverSecs * 1000
    End Sub

    Private Sub AEKPlayer_Closing(sender As Object, e As CancelEventArgs) Handles Me.Closing
        AEKView.MediaPlayer = Nothing
    End Sub

    Private Sub AEKList_Click(sender As Object, e As EventArgs) Handles AEKList.Click
        ' AEKView.MediaPlayer.
    End Sub
End Class