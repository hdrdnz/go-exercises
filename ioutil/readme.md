# ioutil ile Dosya Okuma ve Yazma
+ ioutil paketi standart Golang paketleri içerisinde gelir ve dosya işlemleri yapabilmemiz için bize fonksiyonlar sağlar.
## Dosya Okuma
+ <b>ioutil.Readfile("dosya ismi")</b> ile okunacak dosya belirtilir. Bu geriye dosya ve error olmak üzere iki ifade döndürür.
+  Okuma işlemi byte türünde yapıldığı için dosya içerisindeki ifadeleri  <b>string()</b> ile string tipine dönüştürürüz.

## Dosya Yazma
+ <b>[]byte("yazmak istenilen ifade")</b> ile veri alınır.
+ <b>ioutil.WriteFile("dosya.txt", veri, 0644)</b> ile dosya okunur. "0644" dosya okuma iznidir. bu metod geriyesadece error döndürür.

### Linux Dosya İzinleri

+ İzinlerin harf karşılığı aşağıdaki gibidir.
```
r – okuma yetkisi (read) 
w – yazma yetkisi (write)
x – çalıştırma yetkisi (execute)
```

+ Komut çıktısındaki sayısal ifadeler aşağıdaki şekilde oluşur.
```
4 – okuma (read)
2 – yazma (write)
1 – çalıştırma (execute)
0 – yetki alma
```

+ Örneğin "644" ifadesi ;

<b>Dosya sahibi</b>: 4 + 2  = 6 okuma, yazma  iznine sahiptir.

<b>Dosya sahibiyle aynı grup</b>: 4, okuma  iznine sahiptir.

<b>Diğer kullanıcılar</b>: 4 , okuma ve yazma iznine sahiptir.
