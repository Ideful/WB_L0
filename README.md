для запуска требуется:    


    1) удостовериться в наличии сводобных портов 8080, 4223, 5432  
    
    2) установленная утилита Makefile, Docker и nats-streaming-server  
    
    3) запустить make в корневой папке проекта  
    

для изменения частоты публкации данных в канал нужно изменить значение аргумента функции time.Sleep в nats.go/Publish(), 80 строка nats.go

для проверки attack rate нужно скачать утилиту vegeta и команду make vegeta(также нужно скачать)
