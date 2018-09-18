# io.Copy 사용
사용자정의(?) struct method 사용 예시

[io.Copy](https://golang.org/pkg/builtin/#copy) > CopyBuffer > copyBuffer를 보면 src.Read, dst.Write를 사용하는 것을 확인할 수 있다.  
소스는 아래와 같이 지정된 텍스트 파일의 내용을 표시.  
```sh
$ .\bin\concat.exe .\test.txt
Lorem ipsum 어쩌고저쩌고 브리나응;ㄴ햐ㅗㅓ2039ㅈ루ㅏ너우라ㅜ
ㄴㅇ랴ㅐ효7ㅔ9843ㅐㅓ;ㅑㄷㄹㄴㅇ니로ㅓ
bytes 수: 137
```

source from [this](http://gall.dcinside.com/board/view/?id=programming&no=894823&page=1).

