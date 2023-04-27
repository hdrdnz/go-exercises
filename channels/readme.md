 # Channels (Kanallar)
 + Go routine’ler pratiktir ancak tek başlarına birbirleriyle iletişim kurmadan, bağımsız olarak ve tamamlandıkları zaman herhangi bir sinyal göndermeden çalışırlar. Sessizce çalışıp işleri bittiğinde kullandıkları kaynakları geri iade ederler. 
 + Kanallar sayesinde ise go routine’ler birbirileri ile iletişim kurabilir ve birbirlerine sinyaller göndererek senkronize çalışabilirler. 
 + Bir go routine herhangi bir kanala veri gönderebilir veya kanaldan veri alabilir. Bunu yaparken de verinin akış yönünü gösteren ok (<-) kullanılır. Yani bir nevi atama işlemi yapıyoruz. Atama işleminden farkı, kanala atama işlemi yapılana kadar iş parçacığının devam etmemesidir.

