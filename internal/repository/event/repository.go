package repository_event

import (
	"context"
	model_event "start/internal/model/event"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
)

type eventRepository struct {
	db *pgxpool.Pool
}

func New(conn *pgxpool.Pool) *eventRepository {
	return &eventRepository{db: conn}
}

func (e *eventRepository) GetEvent(ctx context.Context, id int) (model_event.Event, error) {
	query := squirrel.Select("id", "header", "description", "date", "content").
		From("events").
		Where(squirrel.Eq{"id": id})
	err := e.db.Ping(context.Background())
	if err != nil {
		return model_event.Event{}, err
	}

	sql, args, err := query.ToSql()
	if err != nil {
		return model_event.Event{}, err
	}
	if err != nil {
		return model_event.Event{}, err
	}

	row := e.db.QueryRow(ctx, sql, args...)
	var event model_event.Event
	if err := row.Scan(&event.ID, &event.Header, &event.Description, &event.Date, &event.Content); err != nil {
		if err.Error() == "no rows in result set" {
			return model_event.Event{}, nil // No event found
		}
		return model_event.Event{}, err // Other error
	}
	// Assuming the content field is a JSON or similar type that can be scanned directly
	// If Content is a struct, you may need to handle serialization/deserialization
	// For example, if Content is a struct, you might need to convert it to JSON
	// before scanning into the database.
	// If Content is a complex type, you may need to implement custom scanning logic.
	// For example, if Content is a struct, you might need to convert it to JSON
	// before inserting into the database.
	if event.Content.Pictures == nil {
		event.Content.Pictures = []string{} // Initialize to empty slice if nil
	}
	if event.Content.Text == "" {
		event.Content.Text = "" // Initialize to empty string if nil
	}
	if event.Content.Text == "" && len(event.Content.Pictures) == 0 {
		return model_event.Event{}, nil // No content found
	}
	// Return the event with its content
	return event, nil
}

func (e *eventRepository) CreateEvent(ctx context.Context, event model_event.Event) (model_event.Event, error) {
	query := squirrel.Insert("events").
		Columns("header", "description", "date", "content").
		Values(event.Header, event.Description, event.Date, event.Content).
		Suffix("RETURNING id")
	err := e.db.Ping(context.Background())
	if err != nil {
		return model_event.Event{}, err
	}

	sql, args, err := query.ToSql()
	if err != nil {
		return model_event.Event{}, err
	}
	if err != nil {
		return model_event.Event{}, err
	}

	row := e.db.QueryRow(ctx, sql, args...)
	var id int
	if err := row.Scan(&id); err != nil {
		return model_event.Event{}, err
	}
	event.ID = id
	// Assuming the content field is a JSON or similar type that can be stored directly
	// You may need to handle serialization/deserialization if it's a complex type
	// For example, if Content is a struct, you might need to convert it to JSON
	// before inserting into the database.
	return model_event.Event{}, nil
}

func (e *eventRepository) DeleteEvent(ctx context.Context, id int) error {
	_, err := e.db.Exec(context.Background(), "DELETE FROM events WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (e *eventRepository) UpdateEvent(ctx context.Context, event model_event.Event) (model_event.Event, error) {
	query := squirrel.Update("events").
		Set("header", event.Header).
		Set("description", event.Description).
		Set("date", event.Date).
		Set("content", event.Content).
		Where(squirrel.Eq{"id": event.ID}).
		Suffix("RETURNING id")

	err := e.db.Ping(context.Background())
	if err != nil {
		return model_event.Event{}, err
	}

	sql, args, err := query.ToSql()
	if err != nil {
		return model_event.Event{}, err
	}
	if err != nil {
		return model_event.Event{}, err
	}

	row := e.db.QueryRow(ctx, sql, args...)
	var id int
	if err := row.Scan(&id); err != nil {
		return model_event.Event{}, err
	}
	event.ID = id
	return event, nil
}
