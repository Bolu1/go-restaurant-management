package helpers

import(
	"time"
	"github.com/"
)

var UserCollection *mongo.Collection = database.openCollection(database.Client, "user")

var SECRET_KEY string = os.Getenv("SECRET_KEY")

type SignedDetails struct{
	Email string
	First_name string
	Last_name string
	Uid string
	jwt.StandardClaims
}

func GenerateAllToken(email string, firstName string, lastName string, uid string)(signedToken string, signedRefreshToken string, err error){
	claims := &SignedDetails{
		Email: email,
		First_name: firstName,
		Last_name: lastName,
		Uid: uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiredAt: time.Now().Local().Add(time.Hour * time.Duration(358)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SiginMethod256, claims).SignString([]byte(SECRET_KEY))
	refreshtoken, err := jwt.NewWithClaims(jwt.SiginMethod256, refreshClaims).SignedString([]byte(SECRET_KEY))	

	if err != nil{
		log.Fatal(err)
		return 
	}

	return token, refreshToken, err 
}

func UpdateAllToken(signedToken string, signedRefreshToken string, userId string){

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var updatedObj primitive.D

	updatedobj = append(updateObj, bson.E{"token", signedToken})
	updatedobj = append(updateObj, bson.E{"refresh_token", signedRefreshToken})

	Updated_at, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updatedObj = append(updateObj, bson.E{"updated_at", Updatd_at})

	upsert := true
	filter := bson.M{"user_id":userId}
	opt := options.UpdateOptions{
		upsert: &upsert
	}

	_, err := userCollection.UpdateOne{
		ctx,
		filter,
		bson.D{
			{"$set", updateObj},
		},
		&opt,
	}

	defer cancel()

	if err != nil{
		log.Fatal(err)
		return
	}
	return 
}

func ValidateToken(signedTken string)(claims *SignedDetails, ms string){

	token, err = jwt.ParseWithClaims{
		signedToken,
		&SignedDetails(),
		func(token *Jwt.token)(interface[], error){
			return []byte(SECRET_KEY), nil
		}
	}

	claims, ok := token.Claims.(*SignedDetails)
	if !=ok{
		msg = fmt.Sprintf("The token invalid")
		msg = err.Error()
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix(){
		msg  = fmt.Sprint("Token is required")
		msg = err.Error()
		return
	}

	return claims, msg

}