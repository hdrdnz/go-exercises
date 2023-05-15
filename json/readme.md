# JSON Parsing (Ayrıştırma)

+ Günümüzde bir API (application programming interface) a veri göndermede ya da veri çekmede en sık kullanılan veri formatı JSON (javascript object notation) dur

### MARSHALLING (Sıralama)
+ Go programında Go struct’ını JSON stringine dönüştürmek için “encoding” altındaki “json” paketini kullanıyoruz. 

+ <b> json.Marshal("dönüşecek yapı")</b> metodu ile go tipindeki veri JSON tipine dönüşür. veri ve hata olmak üzere iki ifade döndürür. veri byte tipindedir

+ Bildiğimiz gibi Golang’de export etmek için değişken ismi büyük harfle yazılmalıdır. Bu yüzden json tipine çevireceğimiz için elemanları büyük harfte yazmayı unutma.

+ <b>json.MarshalIndent() fonksiyonu</b> json tipine dönüştürelecek verinin satır satır ekrana bastırır.İdentation(boşluk) ekler.
+ json.MarshalIndent("json tipine dönüşecek veri","baş kısmına gelecek kısım","identation")

### UNMARSHALL
+ Unmarshal işlemi amaç olarak marshal işleminin tam tersidir.
+ JSON formatındaki veriyi go struct'ına dönüştürür.


