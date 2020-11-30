<Global.Microsoft.VisualBasic.CompilerServices.DesignerGenerated()>
Partial Class AEKPlayer
    Inherits System.Windows.Forms.Form

    'Form 重写 Dispose，以清理组件列表。
    <System.Diagnostics.DebuggerNonUserCode()>
    Protected Overrides Sub Dispose(ByVal disposing As Boolean)
        Try
            If disposing AndAlso components IsNot Nothing Then
                components.Dispose()
            End If
        Finally
            MyBase.Dispose(disposing)
        End Try
    End Sub

    'Windows 窗体设计器所必需的
    Private components As System.ComponentModel.IContainer

    '注意: 以下过程是 Windows 窗体设计器所必需的
    '可以使用 Windows 窗体设计器修改它。  
    '不要使用代码编辑器修改它。
    <System.Diagnostics.DebuggerStepThrough()>
    Private Sub InitializeComponent()
        Me.components = New System.ComponentModel.Container()
        Dim resources As System.ComponentModel.ComponentResourceManager = New System.ComponentModel.ComponentResourceManager(GetType(AEKPlayer))
        Me.AEKPanel = New System.Windows.Forms.Panel()
        Me.TrackBar1 = New System.Windows.Forms.TrackBar()
        Me.Label1 = New System.Windows.Forms.Label()
        Me.AEKFull = New System.Windows.Forms.Label()
        Me.AEKCurrent = New System.Windows.Forms.Label()
        Me.AEKBar = New System.Windows.Forms.ToolStrip()
        Me.AEKPlay = New System.Windows.Forms.ToolStripButton()
        Me.AEKStop = New System.Windows.Forms.ToolStripButton()
        Me.AEKScreen = New System.Windows.Forms.ToolStripButton()
        Me.AEKList = New System.Windows.Forms.ToolStripButton()
        Me.ToolStripSeparator1 = New System.Windows.Forms.ToolStripSeparator()
        Me.ToolStripButton1 = New System.Windows.Forms.ToolStripDropDownButton()
        Me.设置ToolStripMenuItem = New System.Windows.Forms.ToolStripMenuItem()
        Me.关于ToolStripMenuItem = New System.Windows.Forms.ToolStripMenuItem()
        Me.VLCPanel = New System.Windows.Forms.Panel()
        Me.AEKView = New LibVLCSharp.WinForms.VideoView()
        Me.Timer1 = New System.Windows.Forms.Timer(Me.components)
        Me.TimerBar = New System.Windows.Forms.Timer(Me.components)
        Me.ImageList1 = New System.Windows.Forms.ImageList(Me.components)
        Me.ToolTip1 = New System.Windows.Forms.ToolTip(Me.components)
        Me.AEKPanel.SuspendLayout()
        CType(Me.TrackBar1, System.ComponentModel.ISupportInitialize).BeginInit()
        Me.AEKBar.SuspendLayout()
        Me.VLCPanel.SuspendLayout()
        CType(Me.AEKView, System.ComponentModel.ISupportInitialize).BeginInit()
        Me.SuspendLayout()
        '
        'AEKPanel
        '
        Me.AEKPanel.AllowDrop = True
        Me.AEKPanel.Controls.Add(Me.TrackBar1)
        Me.AEKPanel.Controls.Add(Me.Label1)
        Me.AEKPanel.Controls.Add(Me.AEKFull)
        Me.AEKPanel.Controls.Add(Me.AEKCurrent)
        Me.AEKPanel.Controls.Add(Me.AEKBar)
        Me.AEKPanel.Dock = System.Windows.Forms.DockStyle.Bottom
        Me.AEKPanel.Location = New System.Drawing.Point(0, 483)
        Me.AEKPanel.Name = "AEKPanel"
        Me.AEKPanel.Size = New System.Drawing.Size(1114, 44)
        Me.AEKPanel.TabIndex = 0
        '
        'TrackBar1
        '
        Me.TrackBar1.Location = New System.Drawing.Point(270, 3)
        Me.TrackBar1.Maximum = 100
        Me.TrackBar1.Name = "TrackBar1"
        Me.TrackBar1.Size = New System.Drawing.Size(650, 45)
        Me.TrackBar1.TabIndex = 10
        '
        'Label1
        '
        Me.Label1.AutoSize = True
        Me.Label1.Location = New System.Drawing.Point(978, 11)
        Me.Label1.Name = "Label1"
        Me.Label1.Size = New System.Drawing.Size(131, 12)
        Me.Label1.TabIndex = 9
        Me.Label1.Text = "Label1-HiddenPosition"
        Me.Label1.Visible = False
        '
        'AEKFull
        '
        Me.AEKFull.AutoSize = True
        Me.AEKFull.Location = New System.Drawing.Point(921, 11)
        Me.AEKFull.Name = "AEKFull"
        Me.AEKFull.Size = New System.Drawing.Size(35, 12)
        Me.AEKFull.TabIndex = 3
        Me.AEKFull.Text = "00:00"
        '
        'AEKCurrent
        '
        Me.AEKCurrent.AutoSize = True
        Me.AEKCurrent.Location = New System.Drawing.Point(224, 11)
        Me.AEKCurrent.Name = "AEKCurrent"
        Me.AEKCurrent.Size = New System.Drawing.Size(35, 12)
        Me.AEKCurrent.TabIndex = 2
        Me.AEKCurrent.Text = "00:00"
        '
        'AEKBar
        '
        Me.AEKBar.AutoSize = False
        Me.AEKBar.ImageScalingSize = New System.Drawing.Size(32, 32)
        Me.AEKBar.Items.AddRange(New System.Windows.Forms.ToolStripItem() {Me.AEKPlay, Me.AEKStop, Me.AEKScreen, Me.AEKList, Me.ToolStripSeparator1, Me.ToolStripButton1})
        Me.AEKBar.LayoutStyle = System.Windows.Forms.ToolStripLayoutStyle.Flow
        Me.AEKBar.Location = New System.Drawing.Point(0, 0)
        Me.AEKBar.Name = "AEKBar"
        Me.AEKBar.Size = New System.Drawing.Size(1114, 39)
        Me.AEKBar.TabIndex = 8
        Me.AEKBar.Text = "Navigation"
        '
        'AEKPlay
        '
        Me.AEKPlay.DisplayStyle = System.Windows.Forms.ToolStripItemDisplayStyle.Image
        Me.AEKPlay.Image = CType(resources.GetObject("AEKPlay.Image"), System.Drawing.Image)
        Me.AEKPlay.ImageTransparentColor = System.Drawing.Color.Magenta
        Me.AEKPlay.Name = "AEKPlay"
        Me.AEKPlay.Size = New System.Drawing.Size(36, 36)
        Me.AEKPlay.Text = "ToolStripButton5"
        Me.AEKPlay.ToolTipText = "Play/Open"
        '
        'AEKStop
        '
        Me.AEKStop.DisplayStyle = System.Windows.Forms.ToolStripItemDisplayStyle.Image
        Me.AEKStop.Image = CType(resources.GetObject("AEKStop.Image"), System.Drawing.Image)
        Me.AEKStop.ImageTransparentColor = System.Drawing.Color.Magenta
        Me.AEKStop.Name = "AEKStop"
        Me.AEKStop.Size = New System.Drawing.Size(36, 36)
        Me.AEKStop.Text = "Stop"
        '
        'AEKScreen
        '
        Me.AEKScreen.DisplayStyle = System.Windows.Forms.ToolStripItemDisplayStyle.Image
        Me.AEKScreen.Image = CType(resources.GetObject("AEKScreen.Image"), System.Drawing.Image)
        Me.AEKScreen.ImageTransparentColor = System.Drawing.Color.Magenta
        Me.AEKScreen.Name = "AEKScreen"
        Me.AEKScreen.Size = New System.Drawing.Size(36, 36)
        Me.AEKScreen.Text = "FullScreen"
        '
        'AEKList
        '
        Me.AEKList.DisplayStyle = System.Windows.Forms.ToolStripItemDisplayStyle.Image
        Me.AEKList.Image = CType(resources.GetObject("AEKList.Image"), System.Drawing.Image)
        Me.AEKList.ImageTransparentColor = System.Drawing.Color.Magenta
        Me.AEKList.Name = "AEKList"
        Me.AEKList.Size = New System.Drawing.Size(36, 36)
        Me.AEKList.Text = "Play List"
        Me.AEKList.Visible = False
        '
        'ToolStripSeparator1
        '
        Me.ToolStripSeparator1.AutoSize = False
        Me.ToolStripSeparator1.Name = "ToolStripSeparator1"
        Me.ToolStripSeparator1.Size = New System.Drawing.Size(6, 36)
        '
        'ToolStripButton1
        '
        Me.ToolStripButton1.Alignment = System.Windows.Forms.ToolStripItemAlignment.Right
        Me.ToolStripButton1.DisplayStyle = System.Windows.Forms.ToolStripItemDisplayStyle.Image
        Me.ToolStripButton1.DropDownItems.AddRange(New System.Windows.Forms.ToolStripItem() {Me.设置ToolStripMenuItem, Me.关于ToolStripMenuItem})
        Me.ToolStripButton1.Image = CType(resources.GetObject("ToolStripButton1.Image"), System.Drawing.Image)
        Me.ToolStripButton1.ImageTransparentColor = System.Drawing.Color.Magenta
        Me.ToolStripButton1.Name = "ToolStripButton1"
        Me.ToolStripButton1.RightToLeft = System.Windows.Forms.RightToLeft.No
        Me.ToolStripButton1.Size = New System.Drawing.Size(45, 36)
        Me.ToolStripButton1.Text = "ToolStripButton1"
        Me.ToolStripButton1.Visible = False
        '
        '设置ToolStripMenuItem
        '
        Me.设置ToolStripMenuItem.Name = "设置ToolStripMenuItem"
        Me.设置ToolStripMenuItem.Size = New System.Drawing.Size(100, 22)
        Me.设置ToolStripMenuItem.Text = "设置"
        '
        '关于ToolStripMenuItem
        '
        Me.关于ToolStripMenuItem.Name = "关于ToolStripMenuItem"
        Me.关于ToolStripMenuItem.Size = New System.Drawing.Size(100, 22)
        Me.关于ToolStripMenuItem.Text = "关于"
        '
        'VLCPanel
        '
        Me.VLCPanel.Controls.Add(Me.AEKView)
        Me.VLCPanel.Controls.Add(Me.AEKPanel)
        Me.VLCPanel.Dock = System.Windows.Forms.DockStyle.Fill
        Me.VLCPanel.Location = New System.Drawing.Point(0, 0)
        Me.VLCPanel.Name = "VLCPanel"
        Me.VLCPanel.Size = New System.Drawing.Size(1114, 527)
        Me.VLCPanel.TabIndex = 1
        '
        'AEKView
        '
        Me.AEKView.BackColor = System.Drawing.Color.Black
        Me.AEKView.Dock = System.Windows.Forms.DockStyle.Fill
        Me.AEKView.Location = New System.Drawing.Point(0, 0)
        Me.AEKView.MediaPlayer = Nothing
        Me.AEKView.Name = "AEKView"
        Me.AEKView.Size = New System.Drawing.Size(1114, 483)
        Me.AEKView.TabIndex = 1
        Me.AEKView.Text = "VideoView1"
        '
        'Timer1
        '
        Me.Timer1.Interval = 500
        '
        'TimerBar
        '
        Me.TimerBar.Interval = 8000
        '
        'ImageList1
        '
        Me.ImageList1.ImageStream = CType(resources.GetObject("ImageList1.ImageStream"), System.Windows.Forms.ImageListStreamer)
        Me.ImageList1.TransparentColor = System.Drawing.Color.Transparent
        Me.ImageList1.Images.SetKeyName(0, "gtk-execute.png")
        Me.ImageList1.Images.SetKeyName(1, "gtk-fullscreen.png")
        Me.ImageList1.Images.SetKeyName(2, "kdenlive-zoom-large.png")
        Me.ImageList1.Images.SetKeyName(3, "logo.ico")
        Me.ImageList1.Images.SetKeyName(4, "media-playlist-play.png")
        Me.ImageList1.Images.SetKeyName(5, "player_pause.png")
        Me.ImageList1.Images.SetKeyName(6, "player_play(1).png")
        Me.ImageList1.Images.SetKeyName(7, "player_stop.png")
        Me.ImageList1.Images.SetKeyName(8, "kdenlive-zoom-small.png")
        '
        'AEKPlayer
        '
        Me.AllowDrop = True
        Me.AutoScaleDimensions = New System.Drawing.SizeF(6.0!, 12.0!)
        Me.AutoScaleMode = System.Windows.Forms.AutoScaleMode.Font
        Me.ClientSize = New System.Drawing.Size(1114, 527)
        Me.Controls.Add(Me.VLCPanel)
        Me.Icon = CType(resources.GetObject("$this.Icon"), System.Drawing.Icon)
        Me.Name = "AEKPlayer"
        Me.Text = "AEKPlayer"
        Me.TopMost = True
        Me.AEKPanel.ResumeLayout(False)
        Me.AEKPanel.PerformLayout()
        CType(Me.TrackBar1, System.ComponentModel.ISupportInitialize).EndInit()
        Me.AEKBar.ResumeLayout(False)
        Me.AEKBar.PerformLayout()
        Me.VLCPanel.ResumeLayout(False)
        CType(Me.AEKView, System.ComponentModel.ISupportInitialize).EndInit()
        Me.ResumeLayout(False)

    End Sub

    Friend WithEvents AEKPanel As Panel
    Friend WithEvents AEKFull As Label
    Friend WithEvents AEKCurrent As Label
    Friend WithEvents AEKBar As ToolStrip
    Friend WithEvents AEKPlay As ToolStripButton
    Friend WithEvents AEKStop As ToolStripButton
    Friend WithEvents AEKScreen As ToolStripButton
    Friend WithEvents AEKList As ToolStripButton
    Friend WithEvents ToolStripSeparator1 As ToolStripSeparator
    Friend WithEvents VLCPanel As Panel
    Friend WithEvents ToolStripButton1 As ToolStripDropDownButton
    Friend WithEvents 设置ToolStripMenuItem As ToolStripMenuItem
    Friend WithEvents 关于ToolStripMenuItem As ToolStripMenuItem
    Friend WithEvents Label1 As Label
    Friend WithEvents AEKView As LibVLCSharp.WinForms.VideoView
    Friend WithEvents Timer1 As Timer
    Friend WithEvents TimerBar As Timer
    Friend WithEvents ImageList1 As ImageList
    Friend WithEvents ToolTip1 As ToolTip
    Friend WithEvents TrackBar1 As TrackBar
End Class
