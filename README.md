## Route untuk pengguna (`/users`)
- `POST /register`: Register pengguna baru
- `POST /login`: Login pengguna

## Route untuk foto (`/photos`)
- `POST /`: Membuat foto baru (memerlukan autentikasi)
- `GET /`: Mendapatkan daftar foto
- `GET /:ID`: Mendapatkan foto dengan ID tertentu
- `PUT /:ID`: Mengupdate foto dengan ID tertentu (memerlukan otorisasi dan autentikasi)
- `DELETE /:ID`: Menghapus foto dengan ID tertentu (memerlukan otorisasi dan autentikasi)
- `GET /:ID/comments`: Mendapatkan komentar untuk foto dengan ID tertentu (memerlukan otorisasi)

## Route untuk media sosial (`/social-media`)
- `POST /`: Membuat media sosial baru (memerlukan autentikasi)
- `GET /`: Mendapatkan daftar media sosial
- `GET /:ID`: Mendapatkan media sosial dengan ID tertentu
- `PUT /:ID`: Mengupdate media sosial dengan ID tertentu (memerlukan otorisasi dan autentikasi)
- `DELETE /:ID`: Menghapus media sosial dengan ID tertentu (memerlukan otorisasi dan autentikasi)

## Route untuk komentar (`/comments`)
- `POST /:photoID`: Membuat komentar baru untuk foto dengan ID tertentu (memerlukan autentikasi)
- `GET /:ID`: Mendapatkan komentar dengan ID tertentu
- `PUT /:ID`: Mengupdate komentar dengan ID tertentu (memerlukan otorisasi dan autentikasi)
- `DELETE /:ID`: Menghapus komentar dengan ID tertentu (memerlukan otorisasi dan autentikasi)
