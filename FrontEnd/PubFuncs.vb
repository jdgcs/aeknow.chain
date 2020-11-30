Public Class PubFuncs
    Public Function ZToFSAll(ByVal ZT As Long) As String
        '全存样式
        '从天，bai时duzhi，分，秒dao整合为秒
        Dim T， S， F， M As Integer '天，时，分，秒
        Dim T1， S1 As Integer '天，时，分，秒
        Dim M1 As String
        T = ZT \ 86400
        T1 = ZT Mod 86400
        S = T1 \ 3600
        S1 = T1 Mod 3600
        F = S1 \ 60
        M = S1 Mod 60
        If M < 10 Then
            M1 = "0" & M
        Else
            M1 = M
        End If
        ' ZToFSAll = T & "天" & S & "小时" & F & "分钟" & M1 & "秒"
        ZToFSAll = S & ":" & F & ":" & M1
    End Function
    Public Function ZToFSQS(ByVal ZT As Long) As String
        '缺损样式
        '从秒分解为天，时，分，秒
        Dim T， S， F， M As Integer '天，时，分，秒
        Dim T1， S1 As Integer '天，时，分，秒
        Dim T2， S2， F2， M2 As String '天，时，分，秒
        T = ZT \ 86400
        T1 = ZT Mod 86400
        S = T1 \ 3600
        S1 = T1 Mod 3600
        F = S1 \ 60
        M = S1 Mod 60
        If T = 0 Then
            T2 = ""
        Else
            T2 = T & "天"
        End If
        If S = 0 Then
            S2 = ""
        Else
            S2 = S & "小时"
        End If
        If F = 0 Then
            F2 = ""
        Else
            F2 = F & "分钟"
        End If
        M2 = M & "秒"
        ZToFSQS = T2 & S2 & F2 & M2
    End Function
    Public Function FSoZT(ByVal T As Integer， ByVal S As Int16， ByVal F As Int16， ByVal M As Int16) As Long
        '从天，时，分，秒分解为秒
        'T， S， F， M分别为天，时，分，秒
        FSoZT = T * 86400 + S * 3600 + F * 60 + M
    End Function
End Class
