# Go Kodu Açıklaması - Flutter Karşılaştırması

## 1. Package ve Import (Satır 1-3)

```go
package main

import "fmt"
```

**Açıklama:**
- `package main`: Programın giriş noktası (Flutter'daki `main.dart` gibi)
- `import "fmt"`: Formatting/printing için kütüphane (Flutter'daki `print()` gibi)

**Flutter Karşılaştırması:**
```dart
// Flutter
import 'package:flutter/material.dart';

void main() {
  runApp(MyApp());
}
```

---

## 2. Interface Tanımları (Satır 5-12)

```go
type Putter interface {
	Put(id int, value any) error
}

type Storage interface {
	Putter
	Get(id int) (any, error)
}
```

**Açıklama:**
- **Interface**: Bir sözleşme/şablon. Hangi metodların olması gerektiğini belirtir
- `Putter`: Sadece `Put` metodunu tanımlar
- `Storage`: Hem `Putter`'ı içerir (embedding) hem de `Get` metodunu ekler
- `any`: Herhangi bir tip (Flutter'daki `dynamic` gibi)
- `error`: Hata döndürmek için (Flutter'daki `Exception` gibi)

**Flutter Karşılaştırması:**
```dart
// Flutter - Abstract Class (Interface benzeri)
abstract class Putter {
  Future<void> put(int id, dynamic value);
}

abstract class Storage implements Putter {
  Future<dynamic> get(int id);
}
```

**Neden Kullanılır:**
- Farklı implementasyonlar yapabilirsin (ör: `UserStorage`, `FileStorage`, `DatabaseStorage`)
- Dependency Injection için mükemmel
- Test yazarken mock objeler oluşturabilirsin

---

## 3. Struct ve Metodlar (Satır 14-22)

```go
type UserStorage struct{}

func (s *UserStorage) Get(id int) (any, error) {
	return nil, nil
}

func (s *UserStorage) Put(id int, value any) error {
	return nil
}
```

**Açıklama:**
- `UserStorage struct{}`: Boş bir struct (Flutter'daki class gibi)
- `(s *UserStorage)`: Receiver - bu metodlar `UserStorage`'a ait
- `*`: Pointer receiver - struct'ı değiştirebilir
- `(any, error)`: İki değer döndürür (Go'nun özelliği)

**Flutter Karşılaştırması:**
```dart
// Flutter
class UserStorage implements Storage {
  @override
  Future<dynamic> get(int id) async {
    return null;
  }
  
  @override
  Future<void> put(int id, dynamic value) async {
    // implementation
  }
}
```

**Neden Kullanılır:**
- `Storage` interface'ini implement ediyor
- Gerçek uygulamada veritabanı, dosya sistemi vs. kullanabilirsin

---

## 4. Fonksiyon Parametresi Olarak Interface (Satır 24-26)

```go
func updateValue(id int, value any, storage Storage) error {
	return storage.Put(id, value)
}
```

**Açıklama:**
- `storage Storage`: Interface tipinde parametre
- Herhangi bir `Storage` implementasyonu kabul eder
- Dependency Injection örneği

**Flutter Karşılaştırması:**
```dart
// Flutter
Future<void> updateValue(int id, dynamic value, Storage storage) async {
  await storage.put(id, value);
}
```

**Neden Kullanılır:**
- Esneklik: `UserStorage`, `FileStorage` vs. geçebilirsin
- Test edilebilirlik: Mock storage kullanabilirsin
- Loose coupling: Concrete class'a bağlı değilsin

---

## 5. Struct İçinde Interface (Satır 28-30)

```go
type Server struct {
	storage Storage
}
```

**Açıklama:**
- `Server` struct'ı içinde `Storage` interface'i var
- Composition pattern (Flutter'daki dependency injection gibi)

**Flutter Karşılaştırması:**
```dart
// Flutter
class Server {
  final Storage storage;
  
  Server({required this.storage});
}
```

**Neden Kullanılır:**
- Constructor injection pattern
- Test ederken farklı storage'lar inject edebilirsin
- SOLID prensiplerinden Dependency Inversion

---

## 6. Main Fonksiyonu (Satır 32-46)

```go
func main() {
	server := &Server{
		storage: &UserStorage{},
	}

	server.storage.Put(1, "emirhan")
	value, err := server.storage.Get(1)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Value:", value)
}
```

**Açıklama:**
- `:=`: Kısa değişken tanımlama (tip otomatik çıkarılır)
- `&`: Pointer oluşturur (referans)
- `server.storage.Put()`: Metod çağrısı
- `value, err :=`: İki değer döndürür (Go'nun özelliği)
- `if err != nil`: Hata kontrolü (Go'da standart)

**Flutter Karşılaştırması:**
```dart
// Flutter
void main() async {
  final server = Server(
    storage: UserStorage(),
  );
  
  await server.storage.put(1, "emirhan");
  
  try {
    final value = await server.storage.get(1);
    print("Value: $value");
  } catch (e) {
    print("Error: $e");
  }
}
```

**Farklar:**
- Go: Explicit error return (her fonksiyon error döndürebilir)
- Flutter: Exception/throw kullanır
- Go: `nil` kontrolü yaparsın
- Flutter: try-catch kullanırsın

---

## Özet: Neden Bu Pattern?

1. **Interface Segregation**: Küçük, odaklanmış interface'ler
2. **Dependency Injection**: Test edilebilir kod
3. **Polymorphism**: Farklı implementasyonlar kullanabilme
4. **Error Handling**: Explicit error return (Go'nun felsefesi)

**Gerçek Dünya Kullanımı:**
- Web servisleri (farklı database'ler için)
- File operations (local, cloud storage)
- Testing (mock objects)
- Plugin systems

