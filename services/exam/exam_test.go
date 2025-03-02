package exam

// func setupMockDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	gormDB, err := gorm.Open(postgres.New(postgres.Config{
// 		Conn: db,
// 	}), &gorm.Config{})
// 	if err != nil {
// 		t.Errorf("an error '%s' initializing GORM with mock database", err)
// 	}
// 	return gormDB, mock
// }

// func TestCreateQuizChoices(t *testing.T) {
// 	fmt.Println("TESTING")
// 	var (
// 		gormDB, _ = setupMockDB(t)
// 		now, _    = time.Parse(constants.TIME_LAYOUT, time.Now().Format(constants.TIME_LAYOUT))
// 		quizId    = uuid.FromStringOrNil("e1fc7f53-2fb9-442c-9023-61fa0a6ef26c")
// 		ChoiceId  = uuid.FromStringOrNil("755eb88a-4a42-425a-82d5-cf7f7ad83870")
// 	)
// 	var quiz = &models.Quiz{
// 		ID:        &quizId,
// 		Question:  "test",
// 		CreatedAt: now,
// 		UpdatedAt: nil,
// 		CreatedBy: "test",
// 		UpdatedBy: nil,
// 		Choices: []*models.QuizChoice{
// 			{
// 				ID:        &ChoiceId,
// 				QuizID:    &quizId,
// 				Choice:    "test",
// 				Answer:    "test",
// 				IsCorrect: true,
// 				CreatedAt: now,
// 				UpdatedAt: nil,
// 				CreatedBy: "test",
// 				UpdatedBy: nil,
// 			},
// 		},
// 	}
// 	t.Run("success", func(t *testing.T) {
// 		defer func() {
// 			sqlDB, _ := gormDB.DB()
// 			sqlDB.Close()
// 		}()

// 		examService := exam.NewPsqlExamServiceImpl(gormDB)
// 		_, err := examService.CreateQuizChoices(quiz)
// 		if err != nil {
// 			t.Errorf("Error while creating quiz: %s", err)
// 		}
// 		assert.Error(t, nil)

// 	})
// }
