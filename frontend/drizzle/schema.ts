import { pgTable, serial, varchar, text, integer, timestamp } from 'drizzle-orm/pg-core'

export const urls = pgTable('urls', {
  id: serial('id').primaryKey(),
  shortCode: varchar('short_code', { length: 8 }).unique().notNull(),
  originalUrl: text('original_url').notNull(),
  clicks: integer('clicks').default(0),
  createdAt: timestamp('created_at').defaultNow(),
  userId: integer('user_id'),
})
