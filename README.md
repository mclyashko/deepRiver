# deepRiver
Генератор картинок на основе математических функций. Значение пикселя определяется на основе его координат - x и y

# Запуск:
1. Склонировать репозиторий
2. Перейти в папку с проектом
3. go mod tidy (установка зависимости [Govaluate](github.com/Knetic/govaluate))
4. go build
5. Запуск deepRiver с заданными флагами

# Помощь
\./deepRiver --help

# Примеры использования
* ./deepRiver -fileName="test.png" -size=1024 -mathExp="pow(x,2) + pow(y-pow(x, 2/3), 2)-1"

![test](https://user-images.githubusercontent.com/70536793/180436385-8d4a7415-9bea-4594-957f-23dda27ddc3a.png)
* ./deepRiver -fileName="test.png" -size=1024 -mathExp="x*y"

![test](https://user-images.githubusercontent.com/70536793/180437853-e540c77c-6b34-4e59-9bdb-36ff10bad3cf.png)
* ./deepRiver -fileName="test.png" -size=1024 -mathExp="sin(pow(x,2)+pow(y,2))-cos(x)"

![test](https://user-images.githubusercontent.com/70536793/180437954-7a8b51c9-1bc4-4b2e-8ae9-6abbc5d1007b.png)
* ./deepRiver -fileName="test.png" -size=1024 -mathExp="sin(pow(x,2)+pow(y,2))-cos(x*y)"

![test](https://user-images.githubusercontent.com/70536793/180438064-cb188916-0404-43c7-930c-102791d72d54.png)
* ./deepRiver -fileName="test.png" -size=1024 -mathExp="sin(sin(x)+cos(y))-cos(sin(x*y)+cos(x))"

![test](https://user-images.githubusercontent.com/70536793/180438282-3547585f-086a-405f-b729-061883b73f22.png)
* ./deepRiver -fileName="test.png" -size=1024 -mathExp="sin(x\*y)+cos(x\*y)-1"

![test](https://user-images.githubusercontent.com/70536793/180438367-89a879f8-a086-4b55-95d7-0ba13868de74.png)
* ./deepRiver -fileName="test.png" -size=1024 -mathExp="pow(x, 2) + pow(y, 2)"

![test](https://user-images.githubusercontent.com/70536793/180438502-63c0ed1f-a068-428e-b57a-b9d9b4835924.png)
* ./deepRiver -fileName="test.png" -size=1024 -mathExp="sin(x\*y)\*10 + sin(x)\*2 - cos(y)\*2"

![test](https://user-images.githubusercontent.com/70536793/180438592-a85f6dab-70f3-426f-8a11-096187078276.png)
* ./deepRiver -fileName="test.png" -size=1024 -mathExp="sin(x)/cos(y)"

![test](https://user-images.githubusercontent.com/70536793/180438733-ef041a9c-ca67-4b0d-9b9b-f9941d93f855.png)
* ./deepRiver -fileName="test.png" -size=1024 -mathExp="(pow(pow(x-500,2) + pow(y-500,2) + 2\*100\*(x-500), 2)-4\*pow(100, 2)\*(pow(x-500,2) + pow(y-500,2)) >= 0) ? 255 : 0"

![test](https://user-images.githubusercontent.com/70536793/180438857-deb087fe-887c-4aaf-9f3b-0b9230a58bef.png)
* ./deepRiver -fileName="test.png" -size=1024 -mathExp="sin((x-512)\*(y-512)) + cos((x-512)\*(y-512)) - 1 >= 0 ? 255 : 0"

![test](https://user-images.githubusercontent.com/70536793/180438984-461160a1-b05c-49aa-ba36-5c7ae483abe8.png)
* ./deepRiver -fileName="test.png" -size=1024 -mathExp="pow((y-512)/100, 2) - pow((x-512)/100, 2) + pow((x-512)/100,4) <= 0 ? 255 : 0"

![test](https://user-images.githubusercontent.com/70536793/180439076-ae6dce52-df8e-4ade-b138-7e0b24a8a435.png)

# Как задавать математические функции
[Govaluate](github.com/Knetic/govaluate) отвечает за обработку мат. функций. Функции должны быть составлены согласно [мануалу](https://github.com/Knetic/govaluate/blob/master/MANUAL.md). Функция должна возвращать значение типа float64, это значение будет приведено к типу uint8.
