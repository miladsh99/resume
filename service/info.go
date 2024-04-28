package service

//func UserData(w http.ResponseWriter, r *http.Request) {
//	user := &entity.User{
//		ID:    19,
//		Name:  "raafdgdadin",
//		Email: "arsaddfgaain@gamil.com",
//	}
//	fmt.Fprintf(w, "ID: %d\nUsername: %s\nEmail: %s", user.ID, user.Name, user.Email)
//}
//
//func Authenticate(next http.HandlerFunc) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		c, err := r.Cookie("token")
//		if err != nil {
//			if err == http.ErrNoCookie {
//				w.WriteHeader(http.StatusUnauthorized)
//				return
//			}
//			w.WriteHeader(http.StatusBadRequest)
//			return
//		}
//		tknStr := c.Value
//		claims := &Claims{}
//
//		// Parse and validate the token
//		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
//			return jwtKey, nil
//		})
//		if err != nil {
//			if err == jwt.ErrSignatureInvalid {
//				w.WriteHeader(http.StatusUnauthorized)
//				return
//			}
//			w.WriteHeader(http.StatusBadRequest)
//			return
//		}
//		if !tkn.Valid {
//			w.WriteHeader(http.StatusUnauthorized)
//			return
//		}
//
//		next.ServeHTTP(w, r)
//	}
//}
//
//type Claims struct {
//	Username string `json:"username"`
//	jwt.Claims
//}
