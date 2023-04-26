 ## Goroutine
 ### Concurrency
 
  + Tek işlem birimi veya sıralı işleme kullanarak sistemin yanıt süresini azaltmak için kullanılan bir tekniktir.
  + Bir programı çalıştırdığınızda öncelikle belleğe yüklenir ve bir process olarak nitelendirilir. Process’ler kendi içlerinde bir ya da birden fazla Thread barındırabilirler.Thread dediğimiz şey ise, uygulamamızın çalışan bir parçası anlamına gelmektedir.Modern bilgisayarlarda CPU (Central Processing Unit) içerisinde birden fazla core (çekirdek) bulunur. Process içerisinde bulunan bu Thread’ler bir CPU çekirdeği üzerinde çalışırlar. Her bir core aynı anda sadece bir Thread çalıştırabilir.
 +  Concurrency, paylaşılan kaynaklara çok taraflı erişim sağlar ve bir tür iletişim gerektirir. Yararlı bir ilerleme kaydederken bir iş parçacığı üzerinde çalışır, sonra iş parçacığını durdurur ve yararlı bir ilerleme kaydetmediği takdirde farklı bir iş parçacığına geçer.
 +  Aynı anda birden fazla goroutine çalışmaz. Bir goroutine üzerinde çalışır ancak goroutine üzerinde herhangi bir gecikme olduğunda diğer goroutine çağrılır.
 +  Concurrency'nin çalışma mantığını anlamak için günlük hayat örneği olarak ise hepimiz mesai saatimizde bir proje geliştiriyoruz. Bu projeyi geliştirirken Scrum, Kanban vb. yöntemleri kullanıyoruz. Örneğin, siz bir sprint içerisinde bir işe başladınız ve bu iş için bir branch oluşturdunuz. Geliştirmeye devam ederken production ya da test ortamında bir hata oluştu ve o hatayı çözmeniz gerektiğini düşünün. O esnada çalıştığınız branch’de yazdığınız kodları commit’leyip kaydediyorsunuz ve master/hotfix/fix branch’ine geçiş yapıyorsunuz. Yaşanan hatayı çözdükten sonra da tekrar çalıştığınız branch’e dönüp kaldığınız yerden devam ediyorsunuz. Bu örnek hem context-switch hem de concurrency kavramlarını daha net ve açıklayıcı bir şekilde anlamınızı sağlayacaktır diye düşünüyorum.
 ### Parallelism
 + Birden çok işlemci kullanarak hesaplama hızını artırmak amacıyla tasarlanmıştır.
 + Paralellism için birden fazla işlemci ya da bir işlemci üzerinde birden fazla çekirdeğe (core) sahip olmanız gerekli.Tek CPU ve tek çekirdekli bir makinede paralel kod çalıştıramazsınız.



 ### Concurrency ve Parallelism Arasındaki Temel Farklılıklar
  ![](https://techdifferences.com/wp-content/uploads/2017/12/Untitled.jpg)
 + Concurrency, birden çok görevi aynı anda yürütme ve yönetme eylemidir. Öte yandan Parallelism, çeşitli görevleri aynı anda yürütme eylemidir. <br>
 + Parallelism, çok işlemcili bir sistem gibi birden fazla CPU kullanılarak ve bu işlem birimleri veya CPU'lar üzerinde farklı işlemlerin çalıştırılmasıyla elde edilir. Buna karşılık,Concurrency, CPU üzerindeki işlemlerin serpiştirilmesi ve özellikle bağlam değiştirme ile elde edilir.<br>
 + Eşzamanlılık, tek işlem birimi kullanılarak gerçekleştirilebilirken,  Parallelism  durumunda bu mümkün olamaz, birden fazla işlem birimi gerektirir.
 + Concurrency, yazdığınız kod ve projenizin tasarımı ile alakalı bir kavramdır. Paralellism ise uygulamanızın çalışma anı ile alakalıdır. Paralel çalışacak bir kod yazamazsınız, kodunuz bu özelliğe sadece çalışma zamanında sahip olur.Concurrency, çalışma anında gerçekleşecek paralellism’e kapı aralar. 

 ## Golang'da WaitGroup'u Kullanma
  + WaitGroup,birden fazla goroutine'in bitmesini beklemek için kullanılır.
  WaitGroup, 3 yöntemi dışa aktarır.

  | Metot|Anlam  |
| ------------- | ---------------- |
 Add(int)    | WaitGroup sayacını verilen tamsayı değeri kadar arttırır.          |
| Done()     |  WaitGroup sayacını 1 azaltır, onu bir programın sonlandırıldığını belirtmek için kullanacağız.          |
| Wait()      |  Ana fonksiyonu dahili sayacı 0 olana kadar yürütmeyi engeller.          |



